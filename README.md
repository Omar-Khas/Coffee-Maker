# Coffee-Maker
Coffee machine simulator using Go

Welcome to your automatic barista! This is a simple command-line coffee machine simulator implemented in Go.

## Features

- Simulates a coffee machine that can prepare three types of drinks: Espresso, Latte, and Cappuccino.
- Tracks and displays the current stock of water, milk, coffee beans, disposable cups, and money.
- Allows users to buy drinks, refill resources, take money, and view the remaining resources.

## Usage

1. Clone the repository:

    ```bash
    git clone https://github.com/Omar-Khas/Coffee-Maker.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Coffee-Maker
    ```

3. Run the program:

    ```bash
    go run main.go
    ```

4. Follow the on-screen prompts to interact with the coffee machine.

## Drink Options

1. Espresso's ingredients:
   - Water: 250 ml
   - Coffee Beans: 16 g
     
2. Latte's ingredients:
   - Water: 350 ml
   - Milk: 75 ml
   - Coffee Beans: 20 g

3. Cappuccino's ingredients:
   - Water: 200 ml
   - Milk: 100 ml
   - Coffee Beans: 12 g

## Drink Costs

1. Espresso = $ 4
     
2. Latte = $ 7

3. Cappuccino = $ 6

## Actions

- **Buy**: Purchase a drink by choosing the drink number (1, 2, or 3).
- **Fill**: Refill the machine with water, milk, coffee beans, and disposable cups.
- **Take**: Collect the money earned from selling drinks.
- **Remaining**: View the current stock of resources.
- **Exit**: Quit the coffee machine simulator.
