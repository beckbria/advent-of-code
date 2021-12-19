#include "Problems.h"
/*
--- Day 18: Duet ---

You discover a tablet containing some strange assembly code labeled simply "Duet". Rather than bother the sound card with it, 
you decide to run the code yourself. Unfortunately, you don't see any documentation, so you're left to figure out what the 
instructions mean on your own.

It seems like the assembly is meant to operate on a set of registers that are each named with a single letter and that can 
each hold a single integer. You suppose each register should start with a value of 0.

There aren't that many instructions, so it shouldn't be hard to figure out what they do. Here's what you determine:

snd X plays a sound with a frequency equal to the value of X.
set X Y sets register X to the value of Y.
add X Y increases register X by the value of Y.
mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
mod X Y sets register X to the remainder of dividing the value contained in register X by the value of Y (that is, it sets X 
to the result of X modulo Y).
rcv X recovers the frequency of the last sound played, but only when the value of X is not zero. (If it is zero, the command does nothing.)
jgz X Y jumps with an offset of the value of Y, but only if the value of X is greater than zero. (An offset of 2 skips the 
next instruction, an offset of -1 jumps to the previous instruction, and so on.)

Many of the instructions can take either a register (a single letter) or a number. The value of a register is the integer 
it contains; the value of a number is that number.

After each jump instruction, the program continues with the instruction to which the jump jumped. After any other instruction, 
the program continues with the next instruction. Continuing (or jumping) off either end of the program terminates it.

For example:

set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2

The first four instructions set a to 1, add 2 to it, square it, and then set it to itself modulo 5, resulting in a value of 4.
Then, a sound with frequency 4 (the value of a) is played.
After that, a is set to 0, causing the subsequent rcv and jgz instructions to both be skipped (rcv because a is 0, and jgz 
because a is not greater than 0).
Finally, a is set to 1, causing the next jgz instruction to activate, jumping back two instructions to another jump, which 
jumps again to the rcv, which ultimately triggers the recover operation.

At the time the recover operation is executed, the frequency of the last sound played is 4.

What is the value of the recovered frequency (the value of the most recently played sound) the first time a rcv instruction 
is executed with a non-zero value?

--- Part Two ---

As you congratulate yourself for a job well done, you notice that the documentation has been on the back of the tablet this entire time. 
While you actually got most of the instructions correct, there are a few key differences. This assembly code isn't about sound at all - 
it's meant to be run twice at the same time.

Each running copy of the program has its own set of registers and follows the code independently - in fact, the programs don't even 
necessarily run at the same speed. To coordinate, they use the send (snd) and receive (rcv) instructions:

snd X sends the value of X to the other program. These values wait in a queue until that program is ready to receive them. Each program 
has its own message queue, so a program can never receive a message it sent.
rcv X receives the next value and stores it in register X. If no values are in the queue, the program waits for a value to be sent to it. 
Programs do not continue to the next instruction until they have received a value. Values are received in the order they are sent.

Each program also has its own program ID (one 0 and the other 1); the register p should begin with this value.

For example:

snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d

Both programs begin by sending three values to the other. Program 0 sends 1, 2, 0; program 1 sends 1, 2, 1. Then, each program receives 
a value (both 1) and stores it in a, receives another value (both 2) and stores it in b, and then each receives the program ID of the 
other program (program 0 receives 1; program 1 receives 0) and stores it in c. Each program now sees a different value in its own copy 
of register c.

Finally, both programs try to rcv a fourth time, but no data is waiting for either of them, and they reach a deadlock. When this happens, 
both programs terminate.

It should be noted that it would be equally valid for the programs to run at different speeds; for example, program 0 might have sent 
all three values and then stopped at the first rcv before program 1 executed even its first instruction.

Once both of your programs have terminated (regardless of what caused them to do so), how many times did program 1 send a value?
*/
namespace Day18 {

// The registers are represented by single letters in instructions, but stored as 0-26 internally
inline constexpr unsigned int RegisterIndex(char c) { return (c - 'a'); }

enum class OperandType {
    Invalid,
    Integer,
    Register
};

class Operand {
public:
    Operand(OperandType op, int64_t val) : type(op), value(val) {}
    Operand(const std::string& s);
    Operand() {}
    int64_t Value() const { return value; }
    bool IsRegister() const { return type == OperandType::Register; }
protected:
    OperandType type = OperandType::Invalid;
    int64_t value = 0;
};

Operand::Operand(const std::string& s)
{
    // Registers must be a single character
    if ((s.size() == 1) && (s[0] >= 'a') && (s[0] <= 'z')) {
        type = OperandType::Register;
        value = RegisterIndex(s[0]);
    } else {
        type = OperandType::Integer;
        value = std::stoi(s);
    }
}

class RegisterBank {
public:
    int64_t Evaluate(const Operand& op) const;
    int64_t& operator[] (int x) { return m_register.at(x); }

private:
    std::array<int64_t, 26> m_register = {};
};

int64_t RegisterBank::Evaluate(const Operand& op) const
{
    if (op.IsRegister()) {
        return m_register.at(static_cast<int>(op.Value()));
    } else {
        return op.Value();
    }
}

enum class OpCode {
    Invalid,
    Sound,
    Set,
    Add,
    Multiply,
    Modulo,
    Recover,
    JumpIfGreaterThanZero
};

struct Instruction {
    Instruction(const std::string& inst);
    OpCode type;
    std::array<Operand, 2> operand;
};

Instruction::Instruction(const std::string& inst)
{
    auto tokens = Helpers::Tokenize(inst);

    // Parse the type
    static const struct {
        std::string token;
        OpCode opcode;
    } typeLookup[] = {
        { "snd", OpCode::Sound },
        { "set", OpCode::Set },
        { "add", OpCode::Add },
        { "mul", OpCode::Multiply },
        { "mod", OpCode::Modulo },
        { "rcv", OpCode::Recover },
        { "jgz", OpCode::JumpIfGreaterThanZero },
    };
    type = OpCode::Invalid;
    for (const auto& t : typeLookup) {
        if (tokens[0] == t.token) type = t.opcode;
    }
    if (type == OpCode::Invalid) std::cerr << "Unexpected Instruction: " << tokens[0] << std::endl;

    // Parse the opcodes
    operand[0] = Operand(tokens[1]);
    if (tokens.size() > 2) {
        operand[1] = Operand(tokens[2]);
    }
}

class Program {
public:
    Program(std::vector<Instruction>& instructions) : m_instructions(instructions) {}
    bool Complete() const;
    bool Blocked() const { return m_blocked; }
    int64_t Evaluate(const Operand& op) const { return m_registers.Evaluate(op); }
    void RunInstruction();
    void SetSndHandler(std::function<void(int64_t)> handler) { m_sndHandler = handler; }
    void SetRcvHandler(std::function<void(int64_t)> handler) { m_rcvHandler = handler; }

protected:
    void Snd(const Instruction& inst);
    virtual bool Rcv(const Instruction& inst);      // Returns whether the program is now blocked waiting for a receive
    void Set(const Instruction& inst);
    void Add(const Instruction& inst);
    void Mul(const Instruction& inst);
    void Mod(const Instruction& inst);
    void Jgz(const Instruction& inst);

    int64_t m_programCounter = 0;
    bool m_blocked = false;
    std::vector<Instruction> m_instructions;
    RegisterBank m_registers;
    std::function<void(int64_t)> m_sndHandler;
    std::function<void(int64_t)> m_rcvHandler;
};

bool Program::Complete() const
{
    return (m_programCounter < 0) || (m_programCounter >= static_cast<int>(m_instructions.size()));
}

void Program::RunInstruction()
{
    if (!Complete()) {
        m_blocked = false;
        auto& current = m_instructions[static_cast<int>(m_programCounter++)];
        switch (current.type) {
        case OpCode::Sound:
            Snd(current);
            break;
        case OpCode::Set:
            Set(current);
            break;
        case OpCode::Add:
            Add(current);
            break;
        case OpCode::Multiply:
            Mul(current);
            break;
        case OpCode::Modulo:
            Mod(current);
            break;
        case OpCode::Recover:
            m_blocked = Rcv(current);
            // If we're now blocked, we should stay on this instruction until we're unblocked
            if (m_blocked) --m_programCounter;
            break;
        case OpCode::JumpIfGreaterThanZero:
            Jgz(current);
            break;
        }
    }
}

void Program::Set(const Instruction& inst)
{
    m_registers[static_cast<int>(inst.operand[0].Value())] = Evaluate(inst.operand[1]);
}

void Program::Add(const Instruction& inst) {
    const auto reg = static_cast<int>(inst.operand[0].Value());
    m_registers[reg] = m_registers[reg] + Evaluate(inst.operand[1]);
}

void Program::Mul(const Instruction& inst)
{
    const auto reg = static_cast<int>(inst.operand[0].Value());
    m_registers[reg] = m_registers[reg] * Evaluate(inst.operand[1]);
}

void Program::Mod(const Instruction& inst)
{
    const auto reg = static_cast<int>(inst.operand[0].Value());
    m_registers[reg] = m_registers[reg] % Evaluate(inst.operand[1]);
}

void Program::Jgz(const Instruction& inst)
{
    if (Evaluate(inst.operand[0]) > 0) {
        // We've already incremented the PC by one, so don't double count that space
        m_programCounter += (Evaluate(inst.operand[1]) - 1);
    }
}

void Program::Snd(const Instruction& inst)
{
    if (m_sndHandler) {
        m_sndHandler(Evaluate(inst.operand[0]));
    }
}

bool Program::Rcv(const Instruction& inst)
{
    if (m_rcvHandler) {
        m_rcvHandler(Evaluate(inst.operand[0]));
    }
    return false;
}

class ParallelProgram : public Program {
public:
    ParallelProgram(std::vector<Instruction>& instructions, int64_t programId);
    void PostMessage(int64_t value);
protected:
    virtual bool Rcv(const Instruction& inst) override;

    std::queue<int64_t> m_received;
    std::mutex m_receivedLock;  // Access to the received queue across threads
};

ParallelProgram::ParallelProgram(std::vector<Instruction>& instructions, int64_t programId)
    : Program(instructions)
{
    m_registers[RegisterIndex('p')] = programId;
}

bool ParallelProgram::Rcv(const Instruction& inst)
{
    // If we have no input, wait until we do
    std::lock_guard<std::mutex> lock(m_receivedLock);
    if (m_received.empty()) return true;
    m_registers[static_cast<int>(inst.operand[0].Value())] = m_received.front();
    m_received.pop();
    return false;
}

void ParallelProgram::PostMessage(int64_t value)
{
    std::lock_guard<std::mutex> lock(m_receivedLock);
    m_received.push(value);
}

int64_t RecoverFrequency(const std::vector<std::string>& input)
{
    std::vector<Instruction> instructions;
    for (auto &line : input) {
        instructions.emplace_back(line);
    }

    int64_t mostRecentFrequency = -1;
    bool recovered = false;
    Program program(instructions);
    program.SetSndHandler([&mostRecentFrequency](int64_t frequency){ mostRecentFrequency = frequency; });
    program.SetRcvHandler([&recovered](int64_t condition) { recovered = (condition != 0); });

    while (!program.Complete() && !recovered) {
        program.RunInstruction();
    }
    return mostRecentFrequency;
}

int64_t ParallelExecution(const std::vector<std::string>& input)
{
    std::vector<Instruction> instructions;
    for (auto &line : input) {
        instructions.emplace_back(line);
    }

    ParallelProgram program0(instructions, 0);
    ParallelProgram program1(instructions, 1);
    int64_t program1SndCount = 0;
    program0.SetSndHandler([&program1](int64_t value) {
        program1.PostMessage(value);
    });
    program1.SetSndHandler([&program0, &program1SndCount](int64_t value) {
        program0.PostMessage(value);
        ++program1SndCount;
    });

    while ((!program0.Blocked() || !program1.Blocked()) &&      // Stop if we get to deadlock
            (!program0.Complete() || !program1.Complete())) {    // Stop if both programs have finished
        program0.RunInstruction();
        program1.RunInstruction();
    }
        
    return program1SndCount;
}

} // namespace Day18

void Day18Tests()
{
    const std::vector<std::string> input = {
        "set a 1",
        "add a 2",
        "mul a a",
        "mod a 5",
        "snd a",
        "set a 0",
        "rcv a",
        "jgz a -1",
        "set a 1",
        "jgz a -2"
    };
    const int64_t expected = 4;
    const int64_t result = Day18::RecoverFrequency(input);
    if (result != expected) std::cerr << "Test 18A Error: Got " << result << ", Expected " << expected << std::endl;

    const std::vector<std::string> inputB = {
        "snd 1",
        "snd 2",
        "snd p",
        "rcv a",
        "rcv b",
        "rcv c",
        "rcv d",
    };
    const int64_t expectedB = 3;
    const int64_t resultB = Day18::ParallelExecution(inputB);
    if (resultB != expectedB) std::cerr << "Test 18B Error: Got " << resultB << ", Expected " << expectedB << std::endl;
}

void Day18Problems()
{
    std::cout << "Day 18:\n";
    Day18Tests();
    const auto start = std::chrono::steady_clock::now();
    auto input = Helpers::ReadFileLines("input_day18.txt");
    const auto frequency = Day18::RecoverFrequency(input);
    const auto parallel = Day18::ParallelExecution(input);
    const auto end = std::chrono::steady_clock::now();
    std::cout << frequency << std::endl;
    std::cout << parallel << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl << std::endl;
}