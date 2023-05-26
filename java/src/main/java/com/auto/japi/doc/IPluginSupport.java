package com.auto.japi.doc;

import com.auto.japi.doc.parser.ControllerNode;

import java.util.List;

/**
 * an plugin interface, please feel free to  to do what ever you want.
 *
 * @author yeguozhong yedaxia.github.com
 */
public interface IPluginSupport {

     /**
      * a hook method
      *
      * @param controllerNodeList all the api data
      */
     void execute(List<ControllerNode> controllerNodeList);
}
