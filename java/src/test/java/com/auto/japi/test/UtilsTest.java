package com.auto.japi.test;

import org.junit.Test;
import com.auto.japi.doc.utils.FileUtils;
import com.auto.japi.doc.utils.MavenUtils;

import java.io.File;
import java.io.FileInputStream;
import java.util.Collections;
import java.util.List;

public class UtilsTest {

    @Test
    public void Test_parseDependencyFromPom() throws Exception {
        File file = new File("D:\\document\\workspace\\ads\\test\\pom.xml");
        List<String> list = MavenUtils.parseDependencyFromPom(new FileInputStream(file));
        list.forEach(s -> {
            System.out.println(s);
        });
    }

    @Test
    public void Test_getDependencySourceJar() throws Exception {
        File file = new File("D:\\document\\workspace\\ads\\test\\pom.xml");
        List<String> list = MavenUtils.parseDependencyFromPom(new FileInputStream(file));
        List<String> sourceJars = MavenUtils.getDependencySourceJar(list, "D:\\document\\maven_Repo");
        sourceJars.forEach(s -> {
            System.out.println(s);
        });
    }

    @Test
    public void Test_listFiles() throws Exception {
        File file = new File("D:\\api\\test\\test\\lib");
        file.listFiles((File f) -> {
            if (f.isDirectory()) {
                System.out.println(f.getName());
            }
            return true;
        });
    }

    @Test
    public void Test_writeToFile() throws Exception {
        File file = new File("D:\\opt\\temp\\123\\【新版】  ICON-0601.png");
        FileUtils.writeToFile("D:\\opt\\temp\\123\\1234\\【新版】  ICON-0601-copy.png", new FileInputStream(file));
    }

    @Test
    public void Test_uncompressCodeSourceJar() throws Exception {
        MavenUtils.uncompressCodeSourceJar(Collections.singletonList("D:\\document\\maven_Repo\\com\\test-1.0-SNAPSHOT\\test-sources.jar"), "D:\\opt\\temp\\123\\lib\\source");
    }
}
