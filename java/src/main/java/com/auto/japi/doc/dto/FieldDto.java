package com.auto.japi.doc.dto;

import com.auto.japi.doc.parser.FieldNode;

/**
 * class filed 数据输出实体
 *
 * @author freeguys
 * @date 2022/2/14
 */
public class FieldDto {

    private String name;
    private String type;
    private String description;
    private ClassDto childNode; // 表示该field持有的对象类
    private Boolean loopNode = Boolean.FALSE; // 有循环引用的类
    private Boolean notNull = Boolean.FALSE;


    public static FieldDto buildDto(FieldNode node) {
        if (node == null) {
            return null;
        }
        FieldDto dto = new FieldDto();
        dto.name = node.getName();
        dto.type = node.getType();
        dto.description = node.getDescription();
        dto.childNode = ClassDto.buildDto(node.getChildNode());
        dto.loopNode = node.getLoopNode();
        dto.notNull = node.getNotNull();
        return dto;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public ClassDto getChildNode() {
        return childNode;
    }

    public void setChildNode(ClassDto childNode) {
        this.childNode = childNode;
    }

    public Boolean getLoopNode() {
        return loopNode;
    }

    public void setLoopNode(Boolean loopNode) {
        this.loopNode = loopNode;
    }

    public Boolean getNotNull() {
        return notNull;
    }

    public void setNotNull(Boolean notNull) {
        this.notNull = notNull;
    }
}
