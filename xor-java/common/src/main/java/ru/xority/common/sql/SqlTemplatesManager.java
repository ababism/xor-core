package ru.xority.common.sql;

import com.zaxxer.hikari.HikariConfig;
import com.zaxxer.hikari.HikariDataSource;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.datasource.DataSourceTransactionManager;
import org.springframework.transaction.support.TransactionTemplate;

/**
 * @author foxleren
 */
public class SqlTemplatesManager {
    private final String host;
    private final String port;
    private final String dbName;
    private final String username;
    private final String password;
    private JdbcTemplate jdbcTemplate;
    private TransactionTemplate transactionTemplate;
    private volatile boolean isConfigured;

    public SqlTemplatesManager(
            String host,
            String port,
            String dbName,
            String username,
            String password)
    {
        this.host = host;
        this.port = port;
        this.dbName = dbName;
        this.username = username;
        this.password = password;
        this.isConfigured = false;
    }

    public JdbcTemplate getJdbcTemplate() {
        configureOrSkip();
        return jdbcTemplate;
    }

    public TransactionTemplate getTransactionTemplate() {
        configureOrSkip();
        return transactionTemplate;
    }

    private void configureOrSkip() {
        if (!isConfigured) {
            synchronized (this) {
                if (!isConfigured) {
                    HikariDataSource hikariDataSource = createHikariDataSource();
                    jdbcTemplate = new JdbcTemplate(hikariDataSource);
                    transactionTemplate = new TransactionTemplate(new DataSourceTransactionManager(hikariDataSource));
                    isConfigured = true;
                }
            }
        }
    }

    private HikariDataSource createHikariDataSource() {
        HikariConfig config = new HikariConfig();

        config.setDriverClassName("org.postgresql.Driver");
        config.setJdbcUrl("jdbc:postgresql://" + host + ":" + port + "/" + dbName);
        config.setUsername(username);
        config.setPassword(password);

        return new HikariDataSource(config);
    }
}
