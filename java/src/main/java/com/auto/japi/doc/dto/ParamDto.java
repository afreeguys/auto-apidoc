package com.auto.japi.doc.dto;

import com.auto.japi.doc.parser.ParamNode;

/**
 * request 数据输出实体
 *
 * @author freeguys
 * @date 2022/2/14
 */
public class ParamDto {

    private String name;
    private String type;
    private Boolean required = Boolean.FALSE;
    private String description;

    public static ParamDto buildDto(ParamNode node) {
        if (node == null) {
            return null;
        }
        ParamDto dto = new ParamDto();
        dto.setName(node.getName());
        dto.setType(node.getType());
        dto.setRequired(node.getRequired());
        dto.setDescription(node.getDescription());
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

    public Boolean getRequired() {
        return required;
    }

    public void setRequired(Boolean required) {
        this.required = required;
    }

}
