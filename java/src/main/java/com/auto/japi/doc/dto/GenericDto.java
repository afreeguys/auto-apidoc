package com.auto.japi.doc.dto;

import com.auto.japi.doc.parser.GenericNode;

import java.util.ArrayList;
import java.util.List;

/**
 * 范类 数据输出实体
 *
 * @author freeguys
 * @date 2022/2/14
 */
public class GenericDto {

    private String classType; // for source
    private String modelClass; //for reflection
    private String placeholder;
    private List<GenericDto> childGenericNode = new ArrayList<>();

    public static GenericDto buildDto(GenericNode node) {
        if (node == null) {
            return null;
        }
        GenericDto dto = new GenericDto();
        dto.classType = node.getClassType() == null ? null : node.getClassType().toString();
        dto.modelClass = node.getModelClass() == null ? null : node.getModelClass().getName();
        dto.placeholder = node.getPlaceholder();
        List<GenericNode> childGenericNode = node.getChildGenericNode();
        if (childGenericNode != null && childGenericNode.size() > 0){
            List<GenericDto> genericDtos = new ArrayList<>();
            childGenericNode.forEach(genericNode -> genericDtos.add(GenericDto.buildDto(genericNode)));
            dto.childGenericNode = genericDtos;
        }
        return dto;
    }

    public String getClassType() {
        return classType;
    }

    public void setClassType(String classType) {
        this.classType = classType;
    }

    public String getModelClass() {
        return modelClass;
    }

    public void setModelClass(String modelClass) {
        this.modelClass = modelClass;
    }

    public String getPlaceholder() {
        return placeholder;
    }

    public void setPlaceholder(String placeholder) {
        this.placeholder = placeholder;
    }

    public List<GenericDto> getChildGenericNode() {
        return childGenericNode;
    }

    public void setChildGenericNode(List<GenericDto> childGenericNode) {
        this.childGenericNode = childGenericNode;
    }
}
