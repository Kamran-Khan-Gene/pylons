import { Writer, Reader } from 'protobufjs/minimal';
import { DoubleKeyValue, LongKeyValue, StringKeyValue } from '../pylons/item';
import { Coin } from '../cosmos/base/v1beta1/coin';
import { Recipe } from '../pylons/recipe';
export declare const protobufPackage = "Pylonstech.pylons.pylons";
export interface ItemRecord {
    ID: string;
    doubles: DoubleKeyValue[];
    longs: LongKeyValue[];
    strings: StringKeyValue[];
}
export interface Execution {
    creator: string;
    ID: string;
    nodeVersion: string;
    blockHeight: number;
    itemInputs: ItemRecord[];
    coinOutputs: Coin[];
    itemOutputIDs: string[];
    itemModifyOutputIDs: string[];
    /** we need this recipe stored since the referenced recipe could be updated while this execution is pending */
    recipe: Recipe | undefined;
}
export declare const ItemRecord: {
    encode(message: ItemRecord, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ItemRecord;
    fromJSON(object: any): ItemRecord;
    toJSON(message: ItemRecord): unknown;
    fromPartial(object: DeepPartial<ItemRecord>): ItemRecord;
};
export declare const Execution: {
    encode(message: Execution, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Execution;
    fromJSON(object: any): Execution;
    toJSON(message: Execution): unknown;
    fromPartial(object: DeepPartial<Execution>): Execution;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};