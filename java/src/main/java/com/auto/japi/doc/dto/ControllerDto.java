package com.auto.japi.doc.dto;

import com.auto.japi.doc.parser.ControllerNode;
import com.auto.japi.doc.parser.RequestNode;

import java.util.ArrayList;
import java.util.List;

/**
 * controller 数据输出实体
 *
 * @author freeguys
 * @date 2022/2/14
 */
public class ControllerDto {

    private String author;
    private String description;
    private String baseUrl;
    private String className;
    private String packageName;
    private List<RequestDto> requestNodes = new ArrayList<>();
    private String srcFileName;

    public static ControllerDto buildDto(ControllerNode node){
        if (node == null) {
            return null;
        }
        ControllerDto controllerDto = new ControllerDto();
        controllerDto.setAuthor(node.getAuthor());
        controllerDto.setDescription(node.getDescription());
        controllerDto.setBaseUrl(node.getBaseUrl());
        controllerDto.setClassName(node.getClassName());
        controllerDto.setPackageName(node.getPackageName());
        controllerDto.setSrcFileName(node.getSrcFileName());
        List<RequestNode> requestNodes = node.getRequestNodes();
        if (requestNodes == null ||requestNodes.size() == 0){
            return controllerDto;
        }
        ArrayList<RequestDto> requestDtos = new ArrayList<>();
        requestNodes.forEach(requestNode -> requestDtos.add(RequestDto.buildDto(requestNode)));
        controllerDto.setRequestNodes(requestDtos);
        return controllerDto;
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

    public String getBaseUrl() {
        return baseUrl;
    }

    public void setBaseUrl(String baseUrl) {
        this.baseUrl = baseUrl;
    }

    public String getClassName() {
        return className;
    }

    public void setClassName(String className) {
        this.className = className;
    }

    public String getPackageName() {
        return packageName;
    }

    public void setPackageName(String packageName) {
        this.packageName = packageName;
    }

    public List<RequestDto> getRequestNodes() {
        return requestNodes;
    }

    public void setRequestNodes(List<RequestDto> requestNodes) {
        this.requestNodes = requestNodes;
    }

    public String getSrcFileName() {
        return srcFileName;
    }

    public void setSrcFileName(String srcFileName) {
        this.srcFileName = srcFileName;
    }
}
