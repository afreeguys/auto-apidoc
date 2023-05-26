package com.auto.japi.doc.dto;

/**
 * 数据输出实体
 * @author freeguys
 * @date 2022/2/14
 */
public abstract class Dto<T> {

    /**
     * 指定具体node 转换对应的dto
     */
    abstract T buildDto(T t);
}
