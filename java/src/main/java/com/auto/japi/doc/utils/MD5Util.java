package com.auto.japi.doc.utils;

import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

/**
 * Encoding strings by MD5.
 */
public class MD5Util {

    /**
     * 字符串转MD5(默认Unicode编码)
     *
     * @param key
     * @return
     */
    public static String getStringMD5(String key) {
        String value = null;
        MessageDigest currentAlgorithm;
        try {
            currentAlgorithm = MessageDigest.getInstance("MD5");
            currentAlgorithm.reset();
            currentAlgorithm.update(key.getBytes());
            byte[] hash = currentAlgorithm.digest();
            String d = "";
            int usbyte = 0;
            for (int i = 0; i < hash.length; i += 2) {
                usbyte = hash[i] & 0xFF;
                if (usbyte < 16)
                    d += "0" + Integer.toHexString(usbyte);
                else
                    d += Integer.toHexString(usbyte);
                usbyte = hash[i + 1] & 0xFF;
                if (usbyte < 16)
                    d += "0" + Integer.toHexString(usbyte);
                else
                    d += Integer.toHexString(usbyte);
            }
            value = d.trim().toLowerCase();
        } catch (NoSuchAlgorithmException e) {
            System.out.println("MD5 algorithm not available.");
        }
        return value;
    }

    /**
     * 字符串转MD5带字符编码
     *
     * @param key
     * @param charset
     * @return
     */
    public final static String MD5Encoder(String key, String charset) {
        try {
            byte[] btInput = key.getBytes(charset);
            MessageDigest mdInst = MessageDigest.getInstance("MD5");
            mdInst.update(btInput);
            byte[] md = mdInst.digest();
            StringBuffer sb = new StringBuffer();
            for (int i = 0; i < md.length; i++) {
                int val = ((int) md[i]) & 0xff;
                if (val < 16) {
                    sb.append("0");
                }
                sb.append(Integer.toHexString(val));
            }
            return sb.toString();
        } catch (Exception e) {
            return null;
        }
    }

    /**
     * 加密解密算法 执行一次加密，两次解密
     */
    public static String convertMD5(String inStr) {
        char[] a = inStr.toCharArray();
        for (int i = 0; i < a.length; i++) {
            a[i] = (char) (a[i] ^ 't');
        }
        String s = new String(a);
        return s;
    }

    public static void main(String args[]) {
        String s = new String("cs6324836");
        System.out.println("原始：" + s);
        System.out.println("MD5后：" + getStringMD5(s));
        System.out.println("加密的：" + convertMD5(s));
        System.out.println("解密的：" + convertMD5(convertMD5(s)));
    }
}