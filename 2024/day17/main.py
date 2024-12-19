from enum import IntEnum


class OpCode(IntEnum):
    adv = 0
    bxl = 1
    bst = 2
    jnz = 3
    bxc = 4
    out = 5
    bdv = 6
    cdv = 7


def adv(register: list[int], operand: int) -> None:
    register[0] = int(register[0] / 2 ** combo_operand(operand, register))


def bxl(register: list[int], operand: int) -> None:
    register[1] = register[1] ^ operand


def bst(register: list[int], operand: int) -> None:
    register[1] = combo_operand(operand, register) % 8


def jnz(register: list[int], operand: int) -> int | None:
    if register[0] != 0:
        return operand
    return None


def bxc(register: list[int], operand: int) -> None:
    register[1] = register[1] ^ register[2]


def out(register: list[int], operand: int) -> int:
    return combo_operand(operand, register) % 8


def bdv(register: list[int], operand: int) -> None:
    register[1] = int(register[0] / 2 ** combo_operand(operand, register))


def cdv(register: list[int], operand: int) -> None:
    register[2] = int(register[0] / 2 ** combo_operand(operand, register))


def combo_operand(operand: int, register: list[int]) -> int:
    if operand <= 3:
        return operand
    elif operand < 7:
        return register[operand - 4]
    else:
        raise ValueError("Got an unexpected operand")


DISPATCH = {
    OpCode.adv: adv,
    OpCode.bxl: bxl,
    OpCode.bst: bst,
    OpCode.jnz: jnz,
    OpCode.bxc: bxc,
    OpCode.out: out,
    OpCode.bdv: bdv,
    OpCode.cdv: cdv,
}


def main(input_file="sample.txt"):
    input = open(input_file).read().splitlines()

    program = [OpCode(z) for z in list(map(int, input[4].split(":")[1].split(",")))]
    register = [
        int(input[0].split(":")[1]),
        int(input[1].split(":")[1]),
        int(input[2].split(":")[1]),
    ]

    print(f"Part 1: {part1(register, program)}")
    print(f"Part 2: {part2(program)}")


def part1(register: list[int], program: list[OpCode]):
    output = []
    instruction_pointer = 0
    while instruction_pointer < len(program):
        opcode, operand = program[instruction_pointer], program[instruction_pointer + 1]
        if opcode not in [OpCode.jnz, OpCode.out]:
            DISPATCH[opcode](register, operand)
            instruction_pointer += 2
        if opcode == OpCode.jnz:
            jump_maybe = jnz(register, operand)
            if jump_maybe is not None:
                instruction_pointer = jump_maybe
            else:
                instruction_pointer += 2
        if opcode == OpCode.out:
            output.append(out(register, operand))
            instruction_pointer += 2

    return output


import time


def program_early_break(register: list[int], program: list[OpCode], target: list[int]):
    output = []
    instruction_pointer = 0
    while instruction_pointer < len(program):
        opcode, operand = program[instruction_pointer], program[instruction_pointer + 1]
        if opcode not in [OpCode.jnz, OpCode.out]:
            DISPATCH[opcode](register, operand)
            instruction_pointer += 2
        if opcode == OpCode.jnz:
            jump_maybe = jnz(register, operand)
            if jump_maybe is not None:
                instruction_pointer = jump_maybe
            else:
                instruction_pointer += 2
        if opcode == OpCode.out:
            v = out(register, operand)
            if v != target[len(output)]:
                return None
            output.append(v)
            instruction_pointer += 2

    return output


def checkfast(la, lb):
    for i in range(len(la)):
        if la[-i] != lb[-1]:
            return False
    return True


from itertools import permutations


def part2(program):
    # wow
    target = list(map(int, program))

    # ends = [""]

    # while ends:
    #     end = ends.pop(0)
    #     for prefix in range(8):
    #         newend = f"{end}{prefix}"
    #         l = len(newend)
    #         candidate_regA = int(f"0o{newend}", base=8)

    #         output = part1([candidate_regA, 0, 0], program)
    #         print(l, target, output, newend, candidate_regA)
    #         time.sleep(1)
    #         if output[:l] == target[:l]:
    #             ends.append(newend)
    #     print(ends)

    regA = 8 ** (len(target) - 1)
    print(target, part1([regA, 0, 0], program), oct(regA))
    f = open("day17/pattern.csv", "w")
    for i in range(500_000):
        output = "".join(map(str, part1([regA, 0, 0], program)))
        oct_output = int(output, base=8)
        f.write(f"{regA},{oct_output}\n")
        regA += 1
        # print(output, target)
        # nm = 0
        # while (
        #     nm < len(target) and output[nm] == target[nm] and len(output) == len(target)
        # ):
        #     nm += 1
        # if nm > 3:
        #     print("here!")
        #     print(
        #         regA,
        #         output,
        #         target,
        #         nm,
        #         oct(regA),
        #         oct(
        #             regA,
        #         )[-nm:],
        #     )
        #     # time.sleep(1)
        # break
        # if output == target:
        #     break

    return regA


if __name__ == "__main__":
    main()
