package com.auto.japi.doc.codegenerator.provider;

import com.auto.japi.doc.codegenerator.IFieldProvider;

/**
 * Created by user on 2016/12/25.
 */
public class ProviderFactory {

    public static IFieldProvider createProvider(){
        return new DocFieldProvider();
    }
}
