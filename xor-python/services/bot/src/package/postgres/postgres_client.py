import psycopg2.extras
import psycopg2.pool

from src.internal.config.config import PostgresConfig


class PostgresClient:
    def __init__(self, config: PostgresConfig):
        self._config = config
        self._pool = self.create_pool()

    def create_pool(self):
        return psycopg2.pool.ThreadedConnectionPool(minconn=2,
                                                    maxconn=10,
                                                    host=self._config.host,
                                                    dbname=self._config.dbname,
                                                    user=self._config.user,
                                                    password=self._config.password)

    def execute(self, request, params, commit=True):
        conn = self._pool.getconn()
        cursor = conn.cursor()
        try:
            cursor.execute(request, params)
            result = cursor.fetchone()
            if commit:
                conn.commit()
        except Exception as e:
            print(f"failed to execute query: {e}")
            return None

        self._pool.putconn(conn)
        return result

    def execute_many(self, request, params, commit=True):
        conn = self._pool.getconn()
        cursor = conn.cursor(cursor_factory=psycopg2.extras.DictCursor)
        try:
            cursor.execute(request, params)
            result = cursor.fetchall()
            if commit:
                conn.commit()
        except Exception as e:
            print(f"failed to execute query: {e}")
            return None

        self._pool.putconn(conn)
        return result
