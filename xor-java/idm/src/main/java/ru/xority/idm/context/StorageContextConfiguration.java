package ru.xority.idm.context;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import ru.xority.common.sql.SqlTemplatesManager;
import ru.xority.idm.role.RoleDao;
import ru.xority.idm.role.RolePostgresDao;

/**
 * @author foxleren
 */
@Configuration
public class StorageContextConfiguration {
    @Bean
    public SqlTemplatesManager sqlTemplatesManager(
            @Value("${idm.postgresql.host}") String host,
            @Value("${idm.postgresql.port}") String port,
            @Value("${idm.postgresql.db-name}") String dbName,
            @Value("${idm.postgresql.username}") String username,
            @Value("${idm.postgresql.password}") String password)
    {
        return new SqlTemplatesManager(
                host,
                port,
                dbName,
                username,
                password
        );
    }

    @Bean
    public RoleDao roleDao() {
        return new RolePostgresDao();
    }
}
