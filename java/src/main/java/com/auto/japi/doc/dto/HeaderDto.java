package com.auto.japi.doc.dto;

import com.auto.japi.doc.parser.HeaderNode;

/**
 * header数据输出实体
 *
 * @author freeguys
 * @date 2022/2/14
 */
public class HeaderDto {

    private String name;
    private String value;

    public static HeaderDto buildDto(HeaderNode node) {
        if (node == null) {
            return null;
        }
        HeaderDto dto = new HeaderDto();
        dto.name = node.getName();
        dto.value = node.getValue();
        return dto;
    }

    public String getName() {
        return name;
    }

    public String getValue() {
        return value;
    }
}
