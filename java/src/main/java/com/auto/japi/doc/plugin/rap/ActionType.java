package com.auto.japi.doc.plugin.rap;

/**
 * @author yeguozhong yedaxia.github.com
 */
enum  ActionType {

    GET("1"),
    POST("2"),
    PUT("3"),
    DELETE("4");

    public final String type;

    ActionType(String type) {
        this.type = type;
    }
}
