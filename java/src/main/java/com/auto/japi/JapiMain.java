package com.auto.japi;


import com.auto.japi.doc.Docs;
import com.auto.japi.doc.DocsConfig;
import com.auto.japi.doc.LogUtils;
import com.auto.japi.doc.utils.MavenUtils;

import java.io.File;
import java.io.FileInputStream;
import java.io.InputStream;
import java.util.List;

/**
 * 搜狐java api文档 工具
 * 启发于japidocs (https://japidocs.agilestudio.cn/#/?id=getting-started)
 *
 * @author freeguys
 * @date 2022/2/10
 */
public class JapiMain {

    public static void main(String[] args) throws Exception {
        DocsConfig config = new DocsConfig();
        config.setProjectPath(args[0]); // root project path
        config.setProjectName(args[1]); // project name
        config.setApiVersion(args[2]);  // api version
        config.setDocsPath(args[3]); // api docs target path
        config.setMavenRepository(args[4]);//MavenRepository
        start(config);
    }

    public static void start(DocsConfig config) throws Exception {
        if (!config.checkCanStart()) {
            System.exit(1);
        }
        config.setAutoGenerate(Boolean.TRUE);  // auto generate
        configDependence(config);
        Docs.buildJosnDocs(config); // execute to generate
    }

    /**
     * 配置依赖
     */
    public static void configDependence(DocsConfig config) throws Exception {
        //如果ProjectPath的上级目录有pom文件，说明是聚合工程
        String parentPath = config.getProjectPath();

        while (true) {
            parentPath = parentPath.substring(0, parentPath.lastIndexOf(File.separator));
            if (new File(parentPath + File.separator + "pom.xml").exists()) {
                File p = new File(parentPath);
                p.listFiles((File f) -> {
                    if (f.isDirectory()) {
                        File tmp = new File(f.getAbsolutePath() + File.separator + "src" + File.separator + "main" + File.separator + "java");
                        if (tmp.exists()) {
                            config.addDependencySrcPath(tmp.getAbsolutePath());
                        }
                        return true;
                    }
                    return false;
                });
            } else {
                break;
            }
        }

        File lib = new File(config.getDocsPath() + File.separator + "lib");
        if (!lib.exists()) {
            boolean mkdirs = lib.mkdirs();
            if (!mkdirs) {
                LogUtils.error("lib directory mkdir err");
                System.exit(1);
            }
        }
        File dependencySource = new File(config.getDocsPath() + File.separator + "lib" + File.separator + "dependencySource");
        if (!dependencySource.exists()) {
            boolean mkdirs = dependencySource.mkdirs();
            if (!mkdirs) {
                LogUtils.error("lib/dependencySource directory mkdir err");
                System.exit(1);
            }
        }

        InputStream inputStream = new FileInputStream(new File(config.getProjectPath() + File.separator + "pom.xml"));
        List<String> list = MavenUtils.parseDependencyFromPom(inputStream);
        List<String> sourceJar = MavenUtils.getDependencySourceJar(list, config.getMavenRepository());
        if (sourceJar != null && sourceJar.size() > 0) {
            MavenUtils.uncompressCodeSourceJar(sourceJar, dependencySource.getAbsolutePath());
            config.addDependencySrcPath(dependencySource.getAbsolutePath());
        }
    }

}
