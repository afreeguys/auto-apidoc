package com.auto.japi.doc.doc;

import com.auto.japi.doc.parser.ControllerNode;
import com.auto.japi.doc.DocContext;

import java.util.List;

/**
 * JsonDocGenerator
 *
 * @author freeguys
 * @date 2022/2/10
 */
public class JsonDocGenerator extends AbsDocGenerator {

    public JsonDocGenerator() {
        super(DocContext.controllerParser(), new JsonControllerDocBuilder());
    }

    @Override
    void generateIndex(List<ControllerNode> controllerNodeList) {

    }

}
