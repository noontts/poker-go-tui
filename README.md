# Texas Hold'em Poker (Terminal Edition)

A **simplified Texas Hold'em Poker** game built using **Golang** and the **Bubble Tea** framework. This is a fun, interactive game that runs in your terminal, allowing two players to play through the different phases of poker.

---

## 🎯 **Features**
- Two-player Texas Hold'em Poker
- Terminal-based user interface (TUI) with Bubble Tea
- Game phases: **Deal, Flop, Turn, River, Showdown**
- Interactive gameplay with simple keyboard controls
- Extendable logic for betting, folding, and winner determination

---

## 🚀 **Getting Started**

### Prerequisites
- Go version 1.19+ installed. [Download Go](https://go.dev/dl/)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/poker-terminal.git
   cd poker-terminal
   ```

2. Install the Bubble Tea library:
   ```bash
   go get github.com/charmbracelet/bubbletea
   ```

3. Run the game:
   ```bash
   go run main.go
   ```

---

## 🎮 **How to Play**

### Controls
- **`n`** – Proceed to the next phase (Deal, Flop, Turn, River, Showdown)
- **`q`** or **Esc** – Quit the game

### Game Phases
1. **Deal:** Each player receives 2 cards.
2. **Flop:** 3 community cards are dealt.
3. **Turn:** 1 more community card is dealt.
4. **River:** The final community card is dealt.
5. **Showdown:** Compare hands to determine the winner.

---

## 🃏 **Game Rules**

This version follows the basic rules of **Texas Hold'em Poker**:
- Each player receives **2 private cards** (the "hole" cards).
- There are **5 community cards** revealed over 3 phases (flop, turn, river).
- Players try to make the **best 5-card hand** using any combination of their hole cards and community cards.

---

## 🛠️ **Project Structure**

```
poker-terminal/
├── main.go        # Main entry point
└── utils.go       # Optional utility functions (for card handling, etc.)
```

---

## 🧪 **Testing**

1. Run unit tests (if implemented) with:
   ```bash
   go test ./...
   ```

2. Manually play through multiple rounds to ensure:
   - Cards are dealt correctly.
   - Phases advance as expected.
   - User input works properly.

---

## 🤝 **Contributing**

Feel free to open issues or submit pull requests for any improvements! Contributions are welcome.
