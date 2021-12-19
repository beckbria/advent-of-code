#include "2017/lib/Helpers.h"
/*
--- Day 23: Coprocessor Conflagration ---

You decide to head directly to the CPU and fix the printer from there. As you get close, 
you find an experimental coprocessor doing so much work that the local programs are afraid it will halt and catch fire. 
This would cause serious issues for the rest of the computer, so you head in and see what you can do.

The code it's running seems to be a variant of the kind you saw recently on that tablet. 
The general functionality seems very similar, but some of the instructions are different:

    set X Y sets register X to the value of Y.
    sub X Y decreases register X by the value of Y.
    mul X Y sets register X to the result of multiplying the value contained in register X 
    by the value of Y.
    jnz X Y jumps with an offset of the value of Y, but only if the value of X is not zero. 
    (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)

    Only the instructions listed above are used. The eight registers here, named a through h, 
    all start at 0.

The coprocessor is currently set to some kind of debug mode, which allows for testing, but 
prevents it from doing any meaningful work.

If you run the program (your puzzle input), how many times is the mul instruction invoked?

--- Part Two ---

Now, it's time to fix the problem.

The debug mode switch is wired directly to register a. You flip the switch, which makes register 
a now start at 1 when the program is executed.

Immediately, the coprocessor begins to overheat. Whoever wrote this program obviously didn't 
choose a very efficient implementation. You'll need to optimize the program if it has any hope 
of completing before Santa needs that printer working.

The coprocessor's ultimate goal is to determine the final value left in register h once the 
program completes. Technically, if it had that... it wouldn't even need to run the program.

After setting register a to 1, if the program were to run to completion, what value would be 
left in register h?

*/
namespace Day23
{

    // The registers are represented by single letters in instructions, but stored as 0-26 internally
    inline constexpr unsigned int RegisterIndex(char c) { return (c - 'a'); }

    enum class OperandType
    {
        Invalid,
        Integer,
        Register
    };

    class Operand
    {
    public:
        Operand(OperandType op, int64_t val) : type(op), value(val) {}
        Operand(const std::string &s);
        Operand() {}
        int64_t Value() const { return value; }
        bool IsRegister() const { return type == OperandType::Register; }

    protected:
        OperandType type = OperandType::Invalid;
        int64_t value = 0;
    };

    Operand::Operand(const std::string &s)
    {
        // Registers must be a single character
        if ((s.size() == 1) && (s[0] >= 'a') && (s[0] <= 'z'))
        {
            type = OperandType::Register;
            value = RegisterIndex(s[0]);
        }
        else
        {
            type = OperandType::Integer;
            value = std::stoi(s);
        }
    }

    class RegisterBank
    {
    public:
        int64_t Evaluate(const Operand &op) const;
        int64_t &operator[](int x) { return m_register.at(x); }

    private:
        std::array<int64_t, 26> m_register = {};
    };

    int64_t RegisterBank::Evaluate(const Operand &op) const
    {
        if (op.IsRegister())
        {
            return m_register.at(static_cast<int>(op.Value()));
        }
        else
        {
            return op.Value();
        }
    }

    enum class OpCode
    {
        Invalid,
        Set,
        Subtract,
        Multiply,
        JumpIfNotZero
    };

    struct Instruction
    {
        Instruction(const std::string &inst);
        OpCode type;
        std::array<Operand, 2> operand;
    };

    Instruction::Instruction(const std::string &inst)
    {
        auto tokens = Helpers::Tokenize(inst);

        // Parse the type
        static const struct
        {
            std::string token;
            OpCode opcode;
        } typeLookup[] = {
            {"set", OpCode::Set},
            {"sub", OpCode::Subtract},
            {"mul", OpCode::Multiply},
            {"jnz", OpCode::JumpIfNotZero}};
        type = OpCode::Invalid;
        for (const auto &t : typeLookup)
        {
            if (tokens[0] == t.token)
                type = t.opcode;
        }
        if (type == OpCode::Invalid)
            std::cerr << "Unexpected Instruction: " << tokens[0] << std::endl;

        // Parse the opcodes
        operand[0] = Operand(tokens[1]);
        if (tokens.size() > 2)
        {
            operand[1] = Operand(tokens[2]);
        }
    }

    class Program
    {
    public:
        Program(std::vector<Instruction> &instructions) : m_instructions(instructions) {}
        bool Complete() const;
        int64_t Evaluate(const Operand &op) const { return m_registers.Evaluate(op); }
        void RunInstruction();
        int64_t MulInstructionCount() const { return m_mulCount; }
        void SetRegister(char reg, int64_t value) { m_registers[RegisterIndex(reg)] = value; }
        int64_t GetRegister(char reg) { return m_registers[RegisterIndex(reg)]; }

    protected:
        void Set(const Instruction &inst);
        void Sub(const Instruction &inst);
        void Mul(const Instruction &inst);
        void Jnz(const Instruction &inst);

        int64_t m_programCounter = 0;
        int64_t m_mulCount = 0;
        std::vector<Instruction> m_instructions;
        RegisterBank m_registers;
    };

    bool Program::Complete() const
    {
        return (m_programCounter < 0) || (m_programCounter >= static_cast<int>(m_instructions.size()));
    }

    void Program::RunInstruction()
    {
        if (!Complete())
        {
            auto &current = m_instructions[static_cast<int>(m_programCounter++)];
            switch (current.type)
            {
            case OpCode::Set:
                Set(current);
                break;
            case OpCode::Subtract:
                Sub(current);
                break;
            case OpCode::Multiply:
                Mul(current);
                break;
            case OpCode::JumpIfNotZero:
                Jnz(current);
                break;
            }
        }
    }

    void Program::Set(const Instruction &inst)
    {
        m_registers[static_cast<int>(inst.operand[0].Value())] = Evaluate(inst.operand[1]);
    }

    void Program::Sub(const Instruction &inst)
    {
        const auto reg = static_cast<int>(inst.operand[0].Value());
        m_registers[reg] = m_registers[reg] - Evaluate(inst.operand[1]);
    }

    void Program::Mul(const Instruction &inst)
    {
        const auto reg = static_cast<int>(inst.operand[0].Value());
        m_registers[reg] = m_registers[reg] * Evaluate(inst.operand[1]);
        m_mulCount++;
    }

    void Program::Jnz(const Instruction &inst)
    {
        if (Evaluate(inst.operand[0]) != 0)
        {
            // We've already incremented the PC by one, so don't double count that space
            m_programCounter += (Evaluate(inst.operand[1]) - 1);
        }
    }

    int64_t MultiplyCount(const std::vector<std::string> &input)
    {
        std::vector<Instruction> instructions;
        for (auto &line : input)
        {
            instructions.emplace_back(line);
        }

        Program program(instructions);

        while (!program.Complete())
        {
            program.RunInstruction();
        }
        return program.MulInstructionCount();
    }

    int64_t PrimesInRange()
    {
        // The program was counting non-prime numbers in a very inefficient way.  To see disassembly/processing, see Day23.asm
        int64_t h = 0;
        for (unsigned int b = 105700; b <= 122700; b += 17)
        {
            for (unsigned int d = 2; d < b; d++)
            {
                if (b % d == 0)
                {
                    h++;
                    break;
                }
            }
        }
        return h;
    }

} // namespace Day23

void Day23Problems()
{
    std::cout << "Day 23:\n";
    const auto start = std::chrono::steady_clock::now();
    const auto input = Helpers::ReadFileLines("2017/23/input_day23.txt");
    const auto mulCount = Day23::MultiplyCount(input);
    const auto finalValue = Day23::PrimesInRange();
    const auto end = std::chrono::steady_clock::now();
    std::cout << mulCount << std::endl
              << finalValue << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl
              << std::endl;
}

int main()
{
    Day23Problems();
    return 0;
}