package com.auto.japi.doc.doc;

import com.auto.japi.doc.parser.ControllerNode;

import java.io.IOException;

/**
 * an interface of build a controller api docs
 *
 * @author yeguozhong yedaxia.github.com
 */
public interface IControllerDocBuilder {

    /**
     * build api docs and return as string
     *
     * @param controllerNode
     * @return
     */
    String buildDoc(ControllerNode controllerNode) throws IOException;

}
