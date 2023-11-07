package ru.xority.common.config;

import java.util.Optional;

import org.springframework.boot.web.server.WebServerFactoryCustomizer;
import org.springframework.boot.web.servlet.server.ConfigurableServletWebServerFactory;

/**
 * @author foxleren
 */
public class CommonServerManager {
    private Optional<Integer> port;
    private Optional<String> contextPath;

    public CommonServerManager() {
        port = Optional.empty();
        contextPath = Optional.empty();
    }

    public void setPort(int port) {
        this.port = Optional.of(port);
    }

    public void setContextPath(String contextPath) {
        this.contextPath = Optional.of(contextPath);
    }

    public WebServerFactoryCustomizer<ConfigurableServletWebServerFactory> getWebServerFactoryCustomizer() {
        return factory -> {
            port.ifPresent(factory::setPort);
            contextPath.ifPresent(factory::setContextPath);
        };
    }
}
