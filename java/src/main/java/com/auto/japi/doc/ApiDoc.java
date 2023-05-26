package com.auto.japi.doc;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Retention(RetentionPolicy.SOURCE)
@Target({ElementType.TYPE, ElementType.METHOD})
public @interface ApiDoc {

    /**
     * result class
     * @return
     */
	Class<?> value() default Null.class;

    /**
     * result class
     */
	Class<?> result() default Null.class;

    /**
     * string result: for simple result
     * @return
     */
	String stringResult() default "";

    /**
     * request url
     */
	String url() default "";

    /**
     * request method
     */
	String method() default "get";

    final class Null{

    }
}
