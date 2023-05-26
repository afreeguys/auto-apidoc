package com.auto.japi.doc.utils;

import org.xml.sax.Attributes;
import org.xml.sax.SAXException;
import org.xml.sax.helpers.DefaultHandler;
import com.auto.japi.doc.LogUtils;

import javax.xml.parsers.SAXParser;
import javax.xml.parsers.SAXParserFactory;
import java.io.File;
import java.io.InputStream;
import java.util.ArrayList;
import java.util.Enumeration;
import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;
import java.util.zip.ZipEntry;
import java.util.zip.ZipFile;

public class MavenUtils {


    /**
     * 解析pom文件的的依赖路径
     */
    public static List<String> parseDependencyFromPom(InputStream pomFile) {
        List<String> list = new ArrayList<>();
        try {
            SAXParser saxParser = SAXParserFactory.newInstance().newSAXParser();
            saxParser.parse(pomFile, new DefaultHandler() {

                int count = 0;
                String groupId;
                String artifactId;
                String version;
                boolean isGroupIdTag;
                boolean isArtifactIdTag;
                boolean isVersionTag;


                @Override
                public void startElement(String uri, String localName, String qName, Attributes attributes) throws SAXException {
                    switch (qName) {
                        case "groupId":
                            count = 0;
                            isGroupIdTag = true;
                            break;
                        case "artifactId":
                            isArtifactIdTag = true;
                            break;
                        case "version":
                            isVersionTag = true;
                            break;
                    }
                }

                @Override
                public void endElement(String uri, String localName, String qName) throws SAXException {
                    switch (qName) {
                        case "groupId":
                            isGroupIdTag = false;
                            count++;
                            break;
                        case "artifactId":
                            isArtifactIdTag = false;
                            count++;
                            break;
                        case "version":
                            isVersionTag = false;
                            count++;
                            break;
                    }
                    if (count == 3) {
                        String separator;
                        if (File.separator.equals("\\")) {
                            separator = "\\\\";
                        } else {
                            separator = "/";
                        }
                        String builder = File.separator + groupId.replaceAll("\\.", separator) + File.separator +
                                artifactId + File.separator +
                                version;
                        list.add(builder);
                        count = 0;
                    }
                }

                @Override
                public void characters(char[] ch, int start, int length) throws SAXException {
                    if (isGroupIdTag) {
                        groupId = new String(ch, start, length);
                        return;
                    }
                    if (isArtifactIdTag) {
                        artifactId = new String(ch, start, length);
                        return;
                    }
                    if (isVersionTag) {
                        version = new String(ch, start, length);
                        return;
                    }
                    return;
                }
            });
        } catch (Exception ex) {
            LogUtils.error("read pom.xml error", ex);
        }
        return list;
    }

    /**
     * 根据maven地址获取拥有source的jar包
     */
    public static List<String> getDependencySourceJar(List<String> paths, String mavenRepsPath) {
        if (paths == null || paths.size() == 0) {
            return null;
        }
        return paths.stream().map(s -> {
            String separator;
            if (File.separator.equals("\\")) {
                separator = "\\\\";
            } else {
                separator = "/";
            }
            String[] split = s.split(separator);
            String sourcePath = mavenRepsPath + s + File.separator + split[split.length - 2] + "-" + split[split.length - 1] + "-sources.jar";
            if (new File(sourcePath).exists()) {
                return sourcePath;
            }
            return null;
        }).filter(Objects::nonNull).collect(Collectors.toList());
    }

    public static void uncompressCodeSourceJar(List<String> paths, String sourcePath) throws Exception {
        String sp;
        if (sourcePath.endsWith(File.separator)) {
            sp = sourcePath.substring(0, sourcePath.length() - 1);
        } else {
            sp = sourcePath;
        }
        if (paths == null || paths.size() == 0) {
            return;
        }
        paths.forEach(s -> {
            if (s == null || s.length() == 0) {
                return;
            }
            try {
                ZipFile zipFile = new ZipFile(s);
                Enumeration<? extends ZipEntry> entries = zipFile.entries();
                while (entries.hasMoreElements()) {
                    ZipEntry zipEntry = entries.nextElement();
                    String name = zipEntry.getName();
                    if (name.contains("src/main") || name.contains("java")) {
                        InputStream inputStream = zipFile.getInputStream(zipEntry);
                        FileUtils.writeToFile(sp + File.separator + name.substring(name.lastIndexOf("/") + 1), inputStream);
                    }
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        });
    }
}
