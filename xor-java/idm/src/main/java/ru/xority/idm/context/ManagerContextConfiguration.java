package ru.xority.idm.context;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import ru.xority.idm.role.RoleDao;
import ru.xority.idm.role.RoleManager;

/**
 * @author foxleren
 */
@Configuration
public class ManagerContextConfiguration {
    @Bean
    public RoleManager roleManager(RoleDao roleDao) {
        return new RoleManager(roleDao);
    }
}
