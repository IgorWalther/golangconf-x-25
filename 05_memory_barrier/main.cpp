#include <iostream>
#include <atomic>
#include <thread>
#include <string>

std::string value;
std::atomic<bool> done{false};

void setup() {
    value = "Hello, GolangConf-X-2025!";                      // (3)
    done.store(true, std::memory_order_release);              // (4)
}

int main() {
    std::thread t(setup);

    while (!done.load(std::memory_order_acquire)) {           // (1)

    }

    std::cout << value << std::endl;                          // (2)

    t.join();
    return 0;
}
