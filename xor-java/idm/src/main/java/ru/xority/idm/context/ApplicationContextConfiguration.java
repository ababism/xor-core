package ru.xority.idm.context;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.web.server.WebServerFactoryCustomizer;
import org.springframework.boot.web.servlet.server.ConfigurableServletWebServerFactory;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.support.PropertySourcesPlaceholderConfigurer;
import ru.xority.common.config.CommonConfigManager;
import ru.xority.common.config.CommonServerManager;

/**
 * @author foxleren
 */
@Configuration
public class ApplicationContextConfiguration {
    @Bean
    public PropertySourcesPlaceholderConfigurer propertySourcesPlaceholderConfigurer() {
        return CommonConfigManager.getPropertySourcesPlaceholderConfigurer();
    }

    @Bean
    public WebServerFactoryCustomizer<ConfigurableServletWebServerFactory> webServerFactoryCustomizer(
            @Value("${idm.server.port}") int port,
            @Value("${idm.server.context-path}") String contextPath)
    {
        CommonServerManager commonServerManager = new CommonServerManager();
        commonServerManager.setPort(port);
        commonServerManager.setContextPath(contextPath);
        return commonServerManager.getWebServerFactoryCustomizer();
    }
}
