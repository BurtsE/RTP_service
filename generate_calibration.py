import numpy as np
from scipy.stats import truncnorm
from scipy.integrate import quad
from scipy.optimize import bisect
import json
from pathlib import Path

# --- Константы ---
# Параметры распределения для x
X_MIN = 1.0
X_MAX = 10000.0
# Среднее и стандартное отклонение должны соответствовать диапазону [X_MIN, X_MAX].
# Иначе усеченное распределение будет сконцентрировано на одной из границ.
X_MEAN = (X_MIN + X_MAX) / 2  # Центрируем распределение
X_STD_DEV = (X_MAX - X_MIN) / 6  # Эвристика: 99.7% данных в пределах 3 сигм

# Параметры для поиска k
K_MIN = 0.01
K_MAX = 50.0

# --- Настройка распределения ---
a_param = (X_MIN - X_MEAN) / X_STD_DEV
b_param = (X_MAX - X_MEAN) / X_STD_DEV
dist_X = truncnorm(a_param, b_param, loc=X_MEAN, scale=X_STD_DEV)

def expected_rtp(k):
    """Вычисляет ожидаемый RTP для заданного k."""
    def integrand(bet):
        # bet - это ставка, распределенная как dist_X
        p_survive = 1 - ((bet - 1) / (X_MAX - 1))**(1/k)
        return bet * p_survive * dist_X.pdf(bet)
    integral, _ = quad(integrand, X_MIN, X_MAX, limit=500)
    return integral / dist_X.mean()

def find_k(rtp_target):
    # Находит k для целевого RTP методом деления отрезка пополам.
    def f(k):
        return expected_rtp(k) - rtp_target
    try:
        return bisect(f, K_MIN, K_MAX)
    except ValueError:
        # Возвращаем граничное значение, если rtp_target вне диапазона
        return K_MIN if f(K_MIN) > 0 else K_MAX

# Построим таблицу rtp -> k
rtps = np.linspace(0.01, 1.0, 200)
ks = [find_k(r) for r in rtps]

# Сохраним коэффициенты
output_path = Path("./config/calibration.json")
with output_path.open("w") as f:
    calibration_data = {
        "rtps": rtps.tolist(),
        "ks": ks
    }
    json.dump(calibration_data, f, indent=4)

print(f"Raw RTP and K values saved to '{output_path}'")