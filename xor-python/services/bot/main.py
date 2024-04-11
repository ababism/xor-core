from src.bot import Bot
from src.internal.config.config_provider import ConfigProvider

CONFIG_PATH = "./src/config/config.yml"


def main():
    config_provider = ConfigProvider(CONFIG_PATH)
    config = config_provider.parse_config()

    bot = Bot(config)
    bot.build().run()


if __name__ == '__main__':
    main()
