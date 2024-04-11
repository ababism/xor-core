package ru.xority.idm.api.http;

import java.util.List;
import java.util.UUID;

import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import ru.xority.idm.api.http.dto.RoleResponse;
import ru.xority.idm.service.AccountService;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/admin/account")
@RequiredArgsConstructor
public class AccountAdminController {
    private final AccountService accountService;

    @GetMapping("/{uuid}/roles")
    public ResponseEntity<?> list(@PathVariable UUID uuid) {
        List<RoleResponse> roles = accountService.getRoles(uuid)
                .stream()
                .map(RoleResponse::fromRoleEntity)
                .toList();
        return ResponseEntity.ok(roles);
    }
}
