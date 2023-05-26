package com.auto.japi.doc.doc;

import com.auto.japi.doc.parser.ControllerNode;

import java.io.IOException;

/**
 * build html api docs for a controller
 *
 * @author yeguozhong yedaxia.github.com
 */
public class JsonControllerDocBuilder implements IControllerDocBuilder {

    @Override
    public String buildDoc(ControllerNode controllerNode) throws IOException {

        //todo doc结果转化为json输出
        return "json";
    }

}
