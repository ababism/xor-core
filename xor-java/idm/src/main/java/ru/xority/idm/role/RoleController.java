package ru.xority.idm.role;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/role")
public class RoleController {
    private final RoleManager roleManager;

    public RoleController(RoleManager roleManager) {
        this.roleManager = roleManager;
    }
}
