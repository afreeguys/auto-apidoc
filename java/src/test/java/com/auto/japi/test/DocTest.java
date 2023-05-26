package com.auto.japi.test;

import com.alibaba.fastjson.JSONObject;
import org.apache.commons.io.FileUtils;
import org.junit.Test;
import com.auto.japi.JapiMain;
import com.auto.japi.doc.DocContext;
import com.auto.japi.doc.DocsConfig;
import com.auto.japi.doc.dto.ControllerDto;
import com.auto.japi.doc.parser.ControllerNode;
import com.auto.japi.doc.parser.SpringControllerParser;

import java.io.File;
import java.net.URL;
import java.net.URLClassLoader;
import java.util.ArrayList;
import java.util.Enumeration;
import java.util.List;
import java.util.zip.ZipEntry;
import java.util.zip.ZipFile;

/**
 *
 * @author freeguys
 * @date 2022/3/18
 */
public class DocTest {


    @Test
    public void test() throws Exception {
        DocsConfig config = new DocsConfig();
        config.setProjectPath("D:\\document\\workspace\\ads\\test"); // root project path
        config.setProjectName("test"); // project name
        config.setApiVersion("V6");  // api version
        config.setDocsPath("D:\\api\\test\\test"); // api docs target path
        config.setMavenRepository("D:\\document\\maven_Repo");
        JapiMain.start(config);
    }

    @Test
    public void test_controller_juhe() throws Exception {
        DocsConfig config = new DocsConfig();
        config.setProjectPath("D:\\document\\workspace\\ads\\test"); // root project path
        config.setProjectName("test"); // project name
        config.setApiVersion("V6");  // api version
        config.setDocsPath("D:\\api\\test\\test"); // api docs target path
        config.setMavenRepository("D:\\document\\maven_Repo");
        config.setAutoGenerate(Boolean.TRUE);  // auto generate
        JapiMain.configDependence(config);
        DocContext.init(config);
        SpringControllerParser controllerParser = new SpringControllerParser();
        File file = new File("D:\\document\\workspace\\ads\\test\\src\\main\\java\\com\\test\\controller\\ConsultDataController.java");
        ControllerNode controllerNode = controllerParser.parse(file);
        ControllerDto controllerDto = ControllerDto.buildDto(controllerNode);
        System.out.println(JSONObject.toJSONString(controllerDto));
    }

    @Test
    public void test_controller_jar() throws Exception {
        DocsConfig config = new DocsConfig();
        config.setProjectPath("D:\\document\\workspace\\ads\\test"); // root project path
        config.setProjectName("test"); // project name
        config.setApiVersion("V6");  // api version
        config.setDocsPath("D:\\api\\test\\test"); // api docs target path
        config.setMavenRepository("D:\\document\\maven_Repo");
        config.setAutoGenerate(Boolean.TRUE);  // auto generate
        JapiMain.configDependence(config);
        DocContext.init(config);
        SpringControllerParser controllerParser = new SpringControllerParser();
        File file = new File("D:\\document\\workspace\\ads\\test\\controller\\ad\\AdBiddingTypeController.java");
        ControllerNode controllerNode = controllerParser.parse(file);
        ControllerDto controllerDto = ControllerDto.buildDto(controllerNode);
        System.out.println(JSONObject.toJSONString(controllerDto));
    }

    @Test
    public void loadClass() throws Exception {
        List<URL> urlList = new ArrayList<>();
        ZipFile file = new ZipFile("D:\\document\\workspace\\ads\\test\\target\\test.war");
        Enumeration<? extends ZipEntry> entries = file.entries();
        while (entries.hasMoreElements()){
            ZipEntry zipEntry = entries.nextElement();
            int i = zipEntry.getName().indexOf("WEB-INF/lib");
            if (i >= 0){
                String name = zipEntry.getName().substring(i+12);
                if (name.length() <= 0){
                    continue;
                }
                File tmp = new File("D:\\api\\test\\lib\\"+name);
                FileUtils.copyInputStreamToFile(file.getInputStream(zipEntry),tmp);
                URL url = tmp.toURL();
                urlList.add(url);
            }
        }
        URLClassLoader urlClassLoader = new URLClassLoader(urlList.toArray(new URL[0]));
        Class<?> aClass = urlClassLoader.loadClass("com.test.core.vo.MsgVo");
        System.out.println(aClass.toString());
    }


    @Test
    public void parseClass() throws Exception{
        DocsConfig config = new DocsConfig();
        config.setProjectPath("D:\\document\\workspace\\ads\\test"); // root project path
        config.setProjectName("test"); // project name
        config.setApiVersion("V6");  // api version
        config.setDocsPath("D:\\api\\test"); // api docs target path
        config.setAutoGenerate(Boolean.TRUE);  // auto generate
        config.addLib(new URL("file://D:\\api\\test\\test-1.0-SNAPSHOT.jar4377983892280825901.tmp"));
        DocContext.init(config);

        File file = new File("D:\\document\\workspace\\ads\\test\\controller\\internal\\InternalSiteLandingController.java");
        ControllerNode controllerNode = DocContext.controllerParser().parse(file);
        if (controllerNode.getRequestNodes().isEmpty()) {
            return;
        }
        System.out.println(JSONObject.toJSONString(controllerNode));
    }
}
