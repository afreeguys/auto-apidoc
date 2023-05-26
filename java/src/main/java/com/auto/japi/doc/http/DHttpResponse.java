package com.auto.japi.doc.http;

import com.auto.japi.doc.Utils;

import java.io.IOException;
import java.io.InputStream;
import java.util.HashMap;
import java.util.Map;

/**
 * @author yeguozhong yedaxia.github.com
 */
public class DHttpResponse {

    private int code;
    private InputStream stream;
    private Map<String, String> headers = new HashMap<>();

    public int getCode() {
        return code;
    }

    void setCode(int code) {
        this.code = code;
    }

    public InputStream getStream() {
        return stream;
    }

    void setStream(InputStream stream) {
        this.stream = stream;
    }

    public Map<String, String> getHeaders() {
        return headers;
    }

    void setHeaders(Map<String, String> headers) {
        this.headers = headers;
    }

    void addHeader(String key, String value){
        this.headers.put(key, value);
    }

    public String getHeader(String header){
        return headers.get(header);
    }

    public String streamAsString(){
        try{
            return Utils.streamToString(stream);
        }catch (IOException e){
            throw new RuntimeException(e);
        }
    }
}
