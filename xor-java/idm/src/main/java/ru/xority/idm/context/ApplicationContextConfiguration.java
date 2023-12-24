package ru.xority.idm.context;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.web.server.WebServerFactoryCustomizer;
import org.springframework.boot.web.servlet.server.ConfigurableServletWebServerFactory;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.support.PropertySourcesPlaceholderConfigurer;

import ru.xority.common.serviceconfig.XorConfigManager;
import ru.xority.common.serviceconfig.XorServerManager;

/**
 * @author foxleren
 */
@Configuration
public class ApplicationContextConfiguration {
    @Bean
    public PropertySourcesPlaceholderConfigurer propertySourcesPlaceholderConfigurer() {
        return XorConfigManager.getPropertySourcesPlaceholderConfigurer();
    }

    @Bean
    public WebServerFactoryCustomizer<ConfigurableServletWebServerFactory> webServerFactoryCustomizer(
            @Value("${idm.server.port}") int port,
            @Value("${idm.server.context-path}") String contextPath)
    {
        XorServerManager xorServerManager = new XorServerManager();
        xorServerManager.setPort(port);
        xorServerManager.setContextPath(contextPath);
        return xorServerManager.getWebServerFactoryCustomizer();
    }
}
