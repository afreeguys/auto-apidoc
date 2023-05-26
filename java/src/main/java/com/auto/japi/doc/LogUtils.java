package com.auto.japi.doc;

import java.io.IOException;
import java.util.logging.ConsoleHandler;
import java.util.logging.FileHandler;
import java.util.logging.Handler;
import java.util.logging.Level;
import java.util.logging.Logger;

/**
 * a simple logger
 *
 * @author yeguozhong yedaxia.github.com
 */
public class LogUtils {

    private static final Logger LOGGER = Logger.getGlobal();

    static {
        try {
            Logger parent = LOGGER.getParent();
            Handler[] handlers = parent.getHandlers();
            if (handlers != null && handlers.length > 0) {
                for (Handler handler : handlers) {
                    parent.removeHandler(handler);
                }
            }
            ConsoleHandler consoleHandler = new ConsoleHandler();
            consoleHandler.setFormatter(new LogFormatter());
            consoleHandler.setLevel(Level.ALL);
            parent.addHandler(consoleHandler);
            FileHandler fileHandler = new FileHandler(DocContext.getLogFile().getAbsolutePath());
            fileHandler.setFormatter(new LogFormatter());
            fileHandler.setLevel(Level.ALL);
            LOGGER.addHandler(fileHandler);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void info(String message, Object... args) {
        LOGGER.info(String.format(message, args));
    }

    public static void warn(String message, Object... args) {
        LOGGER.warning(String.format(message, args));
    }

    public static void error(String message, Object... args) {
        LOGGER.severe(String.format(message, args));
    }

    public static void error(String message, Throwable e) {
        LOGGER.log(Level.SEVERE, message, e);
    }
}
