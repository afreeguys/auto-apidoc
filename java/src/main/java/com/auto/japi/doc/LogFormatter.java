package com.auto.japi.doc;

import java.io.PrintWriter;
import java.io.StringWriter;
import java.util.Date;
import java.util.logging.Formatter;
import java.util.logging.LogRecord;

public class LogFormatter extends Formatter {

    // format string for printing the log record
    private final String format = "%1$tY-%1$tm-%1$td %1$tk:%1$tM:%1$tS [%2$s]: %3$s%4$s%n";
    private final Date dat = new Date();

    public LogFormatter() {
    }

    @Override
    public synchronized String format(LogRecord record) {
        dat.setTime(record.getMillis());
        String message = formatMessage(record);
        String throwable = "";
        if (record.getThrown() != null) {
            StringWriter sw = new StringWriter();
            PrintWriter pw = new PrintWriter(sw);
            pw.println();
            record.getThrown().printStackTrace(pw);
            pw.close();
            throwable = sw.toString();
        }
        return String.format(format,
                dat,
                record.getLevel().getName(),
                message,
                throwable);
    }
}
