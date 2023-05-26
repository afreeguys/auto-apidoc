package com.auto.japi.doc;

import com.auto.japi.doc.dto.ControllerDto;
import com.auto.japi.doc.parser.ControllerNode;
import com.auto.japi.doc.plugin.rap.RapSupportPlugin;
import com.auto.japi.doc.doc.JsonDocGenerator;

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.List;

/**
 * main entrance
 */
public class Docs {


    /**
     * build json api docs
     */
    public static void buildJosnDocs(DocsConfig config) throws IOException {
        DocContext.init(config);
        JsonDocGenerator docGenerator = new JsonDocGenerator();
        List<ControllerNode> list = docGenerator.getControllerNodeList();
        File file = new File(config.docsPath + File.separator + config.getApiVersion() + File.separator + "controller.json");
        LogUtils.info(file.getAbsolutePath());
        List<ControllerDto> dtos = new ArrayList<>();
        FileOutputStream outputStream = new FileOutputStream(file);
        try {
            list.forEach(controller -> dtos.add(ControllerDto.buildDto(controller)));
            outputStream.write(Utils.toJson(dtos).getBytes(StandardCharsets.UTF_8));
        } finally {
            outputStream.close();
        }
        DocsConfig docsConfig = DocContext.getDocsConfig();
        if (docsConfig.getRapProjectId() != null && docsConfig.getRapHost() != null) {
            IPluginSupport rapPlugin = new RapSupportPlugin();
            rapPlugin.execute(docGenerator.getControllerNodeList());
        }
        for (IPluginSupport plugin : config.getPlugins()) {
            plugin.execute(docGenerator.getControllerNodeList());
        }
    }

    /**
     * build html api docs
     */
    public static void buildHtmlDocs(DocsConfig config) throws IOException {
        DocContext.init(config);
        JsonDocGenerator docGenerator = new JsonDocGenerator();
        DocContext.setControllerNodeList(docGenerator.getControllerNodeList());
        docGenerator.generateDocs();
        CacheUtils.saveControllerNodes(docGenerator.getControllerNodeList());
        DocsConfig docsConfig = DocContext.getDocsConfig();
        if (docsConfig.getRapProjectId() != null && docsConfig.getRapHost() != null) {
            IPluginSupport rapPlugin = new RapSupportPlugin();
            rapPlugin.execute(docGenerator.getControllerNodeList());
        }
        for (IPluginSupport plugin : config.getPlugins()) {
            plugin.execute(docGenerator.getControllerNodeList());
        }
    }

}
