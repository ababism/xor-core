package ru.xority.common.configuration;

import java.util.Optional;

import org.springframework.boot.web.server.WebServerFactoryCustomizer;
import org.springframework.boot.web.servlet.server.ConfigurableServletWebServerFactory;

/**
 * Class allows customizing server configuration
 * @author foxleren
 */
public class XorServerManager {
    private Optional<Integer> port;
    private Optional<String> contextPath;

    public XorServerManager() {
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
