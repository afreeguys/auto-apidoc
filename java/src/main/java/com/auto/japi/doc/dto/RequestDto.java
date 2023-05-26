package com.auto.japi.doc.dto;

import com.alibaba.fastjson.JSONObject;
import com.auto.japi.doc.parser.ClassNode;
import com.auto.japi.doc.parser.HeaderNode;
import com.auto.japi.doc.parser.ParamNode;
import com.auto.japi.doc.parser.RequestNode;
import com.auto.japi.doc.utils.MD5Util;

import java.util.ArrayList;
import java.util.List;

/**
 * request method 数据实体
 *
 * @author freeguys
 * @date 2022/2/14
 */
public class RequestDto {

    private List<String> method = new ArrayList<>();
    private String url;
    private String methodName; //方法名
    private String author;
    private String description; //接口名
    private String supplement; //补充说明，对应方法 @description
    private List<ParamDto> params = new ArrayList<>();
    private List<ClassDto> requestBodys = new ArrayList<>();
    private List<HeaderDto> header = new ArrayList<>();
    private Boolean deprecated = Boolean.FALSE;
    private ClassDto returnDto;
    private String methodMd5;


    public static RequestDto buildDto(RequestNode node) {
        if (node == null) {
            return null;
        }
        RequestDto dto = new RequestDto();
        dto.setMethod(node.getMethod());
        dto.setUrl(node.getUrl());
        dto.setMethodName(node.getMethodName());
        dto.setAuthor(node.getAuthor());
        dto.setDescription(node.getDescription());
        dto.setSupplement(node.getSupplement());
        dto.setDeprecated(node.getDeprecated());
        dto.setReturnDto(ClassDto.buildDto(node.getResponseNode()));
        List<ParamNode> paramNodes = node.getParamNodes();
        if (paramNodes != null && paramNodes.size() > 0) {
            ArrayList<ParamDto> paramDtos = new ArrayList<>();
            paramNodes.forEach(paramNode -> paramDtos.add(ParamDto.buildDto(paramNode)));
            dto.setParams(paramDtos);
        }
        List<ClassNode> requestBodyNodes = node.getRequestBodyNodes();
        if (requestBodyNodes != null && requestBodyNodes.size() > 0) {
            List<ClassDto> classDtos = new ArrayList<>();
            requestBodyNodes.forEach(classNode -> classDtos.add(ClassDto.buildDto(classNode)));
            dto.setRequestBodys(classDtos);
        }
        List<HeaderNode> headerNodes = node.getHeader();
        if (headerNodes != null && headerNodes.size() > 0) {
            ArrayList<HeaderDto> headerDtos = new ArrayList<>();
            headerNodes.forEach(tmpNode -> headerDtos.add(HeaderDto.buildDto(tmpNode)));
            dto.setHeader(headerDtos);
        }
        dto.setMethodMd5(MD5Util.getStringMD5(JSONObject.toJSONString(dto)));
        return dto;
    }

    public String getMethodMd5() {
        return methodMd5;
    }

    public void setMethodMd5(String methodMd5) {
        this.methodMd5 = methodMd5;
    }

    public List<String> getMethod() {
        return method;
    }

    public void setMethod(List<String> method) {
        this.method = method;
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public String getMethodName() {
        return methodName;
    }

    public void setMethodName(String methodName) {
        this.methodName = methodName;
    }

    public String getAuthor() {
        return author;
    }

    public void setAuthor(String author) {
        this.author = author;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getSupplement() {
        return supplement;
    }

    public void setSupplement(String supplement) {
        this.supplement = supplement;
    }

    public List<ParamDto> getParams() {
        return params;
    }

    public void setParams(List<ParamDto> params) {
        this.params = params;
    }

    public List<ClassDto> getRequestBodys() {
        return requestBodys;
    }

    public void setRequestBodys(List<ClassDto> requestBodys) {
        this.requestBodys = requestBodys;
    }

    public List<HeaderDto> getHeader() {
        return header;
    }

    public void setHeader(List<HeaderDto> header) {
        this.header = header;
    }

    public Boolean getDeprecated() {
        return deprecated;
    }

    public void setDeprecated(Boolean deprecated) {
        this.deprecated = deprecated;
    }

    public ClassDto getReturnDto() {
        return returnDto;
    }

    public void setReturnDto(ClassDto returnDto) {
        this.returnDto = returnDto;
    }
}
