package com.auto.japi.doc.codegenerator.provider;


import com.auto.japi.doc.parser.ClassNode;
import com.auto.japi.doc.parser.FieldNode;
import com.auto.japi.doc.Utils;
import com.auto.japi.doc.codegenerator.IFieldProvider;
import com.auto.japi.doc.codegenerator.model.FieldModel;

import java.util.ArrayList;

import java.util.List;

public class DocFieldProvider implements IFieldProvider {

	@Override
	public List<FieldModel> provideFields(ClassNode respNode) {
		List<FieldNode>recordNodes = respNode.getChildNodes();
		if(recordNodes == null || recordNodes.isEmpty()){
			return null;
		}
		List<FieldModel> entryFieldList = new ArrayList<>();
		FieldModel entryField;
		for (FieldNode recordNode : recordNodes) {
			entryField = new FieldModel();
			String fieldName = DocFieldHelper.getPrefFieldName(recordNode.getName());
			entryField.setCaseFieldName(Utils.capitalize(fieldName));
			entryField.setFieldName(fieldName);
			entryField.setFieldType(DocFieldHelper.getPrefFieldType(recordNode.getType()));
			entryField.setRemoteFieldName(recordNode.getName());
			entryField.setComment(recordNode.getDescription());
			entryFieldList.add(entryField);
		}
		return entryFieldList;
	}
	
}
