package com.auto.japi.doc.loader;

import java.io.File;
import java.net.MalformedURLException;
import java.net.URL;
import java.net.URLClassLoader;
import java.util.ArrayList;
import java.util.List;

public class LibClassLoader extends URLClassLoader {

    public LibClassLoader(URL[] urls) {
        super(urls);
    }

    public LibClassLoader(String libPath) {
        super(loadLib(libPath));
    }

    /**
     * @param libHome
     * @return
     */
    private static URL[] loadLib(String libHome) {
        // lib目录下包的URL
        List<URL> urlList = new ArrayList<URL>();
        // 递归
        loopLibFiles(new File(libHome), urlList);
        // 转换
        URL[] urlArray = new URL[urlList.size()];
        for(int i=0;i<urlList.size();i++){
            urlArray[0] = urlList.get(i);
        }
        return urlArray;
    }


    /**
     * 递归
     *
     * @param file
     * @param urlList
     */
    private static final void loopLibFiles(File file, List<URL> urlList) {
        if (file.isDirectory()) {
            File[] tmps = file.listFiles();
            for (File tmp : tmps) {
                loopLibFiles(tmp, urlList);
            }
        }else {
            String p = file.getAbsolutePath();
            if (p.endsWith(".jar")) {
                try {
                    urlList.add(file.toURI().toURL());
                } catch (MalformedURLException e) {
                    e.printStackTrace();
                }
            }
        }
    }
}
