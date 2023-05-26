package com.auto.japi.doc.dto;

import com.auto.japi.doc.parser.ClassNode;
import com.auto.japi.doc.parser.FieldNode;
import com.auto.japi.doc.parser.GenericNode;

import java.util.ArrayList;
import java.util.List;

/**
 * object 对应实体
 *
 * @author freeguys
 * @date 2022/2/14
 */
public class ClassDto {

    private String className = "";
    private String modelClass; //for reflection
    private String description;
    private Boolean isList = Boolean.FALSE;
    private List<FieldDto> fields = new ArrayList<>();
    private List<GenericDto> genericNodes = new ArrayList<>();

    public static ClassDto buildDto(ClassNode node) {
        if (node == null) {
            return null;
        }
        ClassDto dto = new ClassDto();
        dto.className = node.getClassName();
        dto.modelClass = node.getModelClass() == null ? null : node.getModelClass().getName();
        dto.description = node.getDescription();
        dto.description = node.toJsonApi();
        List<FieldNode> childNodes = node.getChildNodes();
        if (childNodes != null && childNodes.size() > 0) {
            List<FieldDto> fieldDtos = new ArrayList<>();
            childNodes.forEach(fieldNode -> fieldDtos.add(FieldDto.buildDto(fieldNode)));
            dto.fields = fieldDtos;
        }
        List<GenericNode> genericNodes = node.getGenericNodes();
        if (genericNodes != null && genericNodes.size() > 0) {
            List<GenericDto> genericDtos = new ArrayList<>();
            genericNodes.forEach(fieldNode -> genericDtos.add(GenericDto.buildDto(fieldNode)));
            dto.genericNodes = genericDtos;
        }
        return dto;
    }

    /**
     * class ParentNode{ //parentNode;
     * ClassNode node;
     * }
     */

    private Boolean showFieldNotNull = Boolean.FALSE;

    public String getClassName() {
        return className;
    }

    public void setClassName(String className) {
        this.className = className;
    }

    public String getModelClass() {
        return modelClass;
    }

    public void setModelClass(String modelClass) {
        this.modelClass = modelClass;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public Boolean getList() {
        return isList;
    }

    public void setList(Boolean list) {
        isList = list;
    }

    public List<FieldDto> getFields() {
        return fields;
    }

    public void setFields(List<FieldDto> fields) {
        this.fields = fields;
    }

    public List<GenericDto> getGenericNodes() {
        return genericNodes;
    }

    public void setGenericNodes(List<GenericDto> genericNodes) {
        this.genericNodes = genericNodes;
    }

    public Boolean getShowFieldNotNull() {
        return showFieldNotNull;
    }

    public void setShowFieldNotNull(Boolean showFieldNotNull) {
        this.showFieldNotNull = showFieldNotNull;
    }
}
