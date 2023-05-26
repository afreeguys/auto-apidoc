package com.auto.japi.doc.parser;

/**
 * Created by lzw on 2017/8/23.
 */
public class HeaderNode {

    private String name;
    private String value;

    public HeaderNode(String name, String value) {
        this.name = name;
        this.value = value;
    }

    public String getName() {
        return name;
    }

    public String getValue() {
        return value;
    }
}
