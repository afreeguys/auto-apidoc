package com.auto.japi.doc;

import com.auto.japi.doc.plugin.rap.RapSupportPlugin;

import java.net.URL;
import java.util.ArrayList;
import java.util.List;
import java.util.Locale;

/**
 * @author yeguozhong yedaxia.github.com
 */
public class DocsConfig {

    String projectPath; // must set
    List<String> javaSrcPaths = new ArrayList<>(); //multi modules support
    List<String> dependencySrcPaths = new ArrayList<>(); //multi modules support
    String docsPath; // default equals projectPath
    String resourcePath; // if empty, use the default resources
    String mvcFramework; //spring, play, jfinal, generic, can be empty
    String apiVersion; // this api version
    String projectName; //project name
    Boolean autoGenerate = Boolean.FALSE; // 自动生成所有Controller的接口文档，不需要@ApiDoc注解
    Locale locale = Locale.getDefault();
    Boolean openReflection = Boolean.TRUE; // 是否开启对象反射
    String mavenRepository; // jar或war包的路径

    String rapHost;
    String rapLoginCookie;
    String rapProjectId;
    String rapAccount;
    String rapPassword;

    List<IPluginSupport> plugins = new ArrayList<>();

    List<URL> libs = new ArrayList<>();

    List<IPluginSupport> getPlugins() {
        return plugins;
    }

    /**
     * add your own plugin, example:
     * @see RapSupportPlugin
     * @param plugin
     */
    public void addPlugin(IPluginSupport plugin) {
        this.plugins.add(plugin);
    }

    boolean isSpringMvcProject(){
        return mvcFramework != null && mvcFramework.equals("spring");
    }

    boolean isPlayProject(){
        return mvcFramework != null && mvcFramework.equals("play");
    }

    boolean isJfinalProject(){
        return mvcFramework != null && mvcFramework.equals("jfinal");
    }

    boolean isGeneric(){
        return mvcFramework != null && mvcFramework.equals("generic");
    }

    public void setProjectPath(String projectPath) {
        this.projectPath = projectPath;
    }

    public String getProjectPath() {
        return projectPath;
    }

    public void setDocsPath(String docsPath) {
        this.docsPath = docsPath;
    }

    public String getDocsPath() {
        return docsPath;
    }

    public void setMvcFramework(String mvcFramework) {
        this.mvcFramework = mvcFramework;
    }

    public List<String> getJavaSrcPaths() {
        return javaSrcPaths;
    }

    public void addDependencySrcPath(String d) {
        dependencySrcPaths.add(d);
    }

    public List<String> getDependencySrcPaths() {
        return dependencySrcPaths;
    }

    public String getApiVersion() {
        return apiVersion;
    }

    public void setApiVersion(String apiVersion) {
        this.apiVersion = apiVersion;
    }

    public Boolean getAutoGenerate() {
        return autoGenerate;
    }

    public void setAutoGenerate(Boolean autoGenerate) {
        this.autoGenerate = autoGenerate;
    }

    public String getProjectName() {
        return projectName;
    }

    public void setProjectName(String projectName) {
        this.projectName = projectName;
    }

    public Locale getLocale() {
        return locale;
    }

    public void setLocale(Locale locale) {
        this.locale = locale;
    }

    public String getResourcePath() {
        return resourcePath;
    }

    public void setResourcePath(String resourcePath) {
        this.resourcePath = resourcePath;
    }

    /**
     * if cannot find the java file from other module, you can try to config the java src path yourself.
     *
     * @param javaSrcPath
     */
    public void addJavaSrcPath(String javaSrcPath){
        javaSrcPaths.add(javaSrcPath);
    }

    public String getRapHost() {
        return rapHost;
    }

    public String getRapLoginCookie() {
        return rapLoginCookie;
    }

    public String getRapProjectId() {
        return rapProjectId;
    }

    public String getRapAccount() {
        return rapAccount;
    }

    public String getRapPassword() {
        return rapPassword;
    }

    public Boolean getOpenReflection() {
        return openReflection;
    }

    public String getMavenRepository() {
        return mavenRepository;
    }

    public void setMavenRepository(String mavenRepository) {
        this.mavenRepository = mavenRepository;
    }

    public void setOpenReflection(Boolean openReflection) {
        this.openReflection = openReflection;
    }

    public void addLib(URL url) {
        if (url == null) {
            return;
        }
        libs.add(url);
    }

    public List<URL> getLibs() {
        return libs;
    }



    /**
     * 校验是否可以启动
     */
    public boolean checkCanStart() {
        if (this.projectPath == null) {
            LogUtils.error("projectPath is null");
            return false;
        }
        if (this.projectName == null) {
            LogUtils.error("projectName is null");
        }
        if (this.apiVersion == null) {
            LogUtils.error("apiVersion is null");
        }
        if (this.docsPath == null) {
            LogUtils.error("docsPath is null");
        }
        if (this.mavenRepository == null) {
            LogUtils.error("mavenRepository is null");
        }
        return true;
    }
}
