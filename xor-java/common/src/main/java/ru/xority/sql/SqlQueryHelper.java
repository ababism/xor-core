package ru.xority.sql;

import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.List;

/**
 * @author foxleren
 */
public class SqlQueryHelper {
    public static void buildPreparedStatement(PreparedStatement ps, List<Object> values) throws SQLException {
        int i = 1;
        for (Object arg : values) {
            ps.setObject(i++, arg);
        }
    }

    public static String queryWhereAnd(List<String> keys) {
        StringBuilder sb = new StringBuilder(" WHERE ");

        int i = 1;
        for (String key : keys) {
            if (i > 1) {
                sb.append(" AND ");
            }
            sb.append(String.format("%s = ?", key));
            i++;
        }
        return sb.toString();
    }
}
