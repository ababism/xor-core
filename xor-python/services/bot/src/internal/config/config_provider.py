import logging

import yaml

from src.internal.config.config import Config

logger = logging.getLogger(__name__)


class ConfigProvider:
    def __init__(self, config_path: str):
        self._config_path = config_path

    def parse_config(self) -> Config:
        try:
            default_config = self._read_default_config()
            return Config.model_validate(default_config)
        except Exception as e:
            logger.error(f"Failed to parse config: {e}")
            raise RuntimeError(e)

    def _read_default_config(self) -> dict:
        with open(self._config_path) as config_file:
            return yaml.safe_load(config_file)
