package com.auto.japi.doc.utils;

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;

public class FileUtils {

    /**
     * 将content写入文件，如果文件存在就不会写入了
     */
    public static void writeToFile(String dist, InputStream content) throws IOException {
        File file = new File(dist);
        if (!file.exists()) {
            int i = dist.lastIndexOf(File.separator);
            File parentDir = new File(dist.substring(0, i));
            if (!parentDir.exists()) {
                if (!parentDir.mkdirs()) {
                    throw new IOException("mkdir dir[" + parentDir.getAbsolutePath() + "] err");
                }
            }
            if (!file.createNewFile()) {
                throw new IOException("create file[" + dist + "] err");
            }
        }
        try (FileOutputStream outputStream = new FileOutputStream(file)) {
            byte[] bytes = new byte[1024];
            while (true) {
                int n = content.read(bytes);
                if (n < 0) {
                    return;
                }
                outputStream.write(bytes, 0, n);
            }
        }


    }
}
