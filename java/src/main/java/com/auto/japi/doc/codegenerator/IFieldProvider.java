package com.auto.japi.doc.codegenerator;

import com.auto.japi.doc.codegenerator.model.FieldModel;
import com.auto.japi.doc.parser.ClassNode;

import java.util.List;

public interface IFieldProvider {
	/**
	 * get response fields
	 * @param respNode
	 * @return
	 */
	List<FieldModel> provideFields(ClassNode respNode);
}
