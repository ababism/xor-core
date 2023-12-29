package ru.xority.common.configuration;

import org.springframework.context.support.PropertySourcesPlaceholderConfigurer;
import org.springframework.core.io.ClassPathResource;
import org.springframework.core.io.Resource;

/**
 * Class allows using properties from common project
 * @author foxleren
 */
public class XorCommonPropertiesProvider {
    private static final String COMMON_PROPERTIES = "common.properties";

    public static PropertySourcesPlaceholderConfigurer getPropertySourcesPlaceholderConfigurer() {
        PropertySourcesPlaceholderConfigurer pspc = new PropertySourcesPlaceholderConfigurer();
        Resource[] resources = new ClassPathResource[]{
                new ClassPathResource(COMMON_PROPERTIES)
        };
        pspc.setLocations(resources);
        return pspc;
    }
}
