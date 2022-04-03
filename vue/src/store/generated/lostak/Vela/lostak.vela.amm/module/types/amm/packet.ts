/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "lostak.vela.amm";

export interface AmmPacketData {
  noData: NoData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field */
  addLiquidityPacket: AddLiquidityPacketData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field/number */
  createPoolPacket: CreatePoolPacketData | undefined;
}

export interface NoData {}

/** CreatePoolPacketData defines a struct for the packet payload */
export interface CreatePoolPacketData {
  poolParam: string;
  poolAssets: string;
}

/** CreatePoolPacketAck defines a struct for the packet acknowledgment */
export interface CreatePoolPacketAck {
  poolId: number;
}

/** AddLiquidityPacketData defines a struct for the packet payload */
export interface AddLiquidityPacketData {
  amountVela: string;
  amountOther: string;
}

/** AddLiquidityPacketAck defines a struct for the packet acknowledgment */
export interface AddLiquidityPacketAck {
  amountShares: string;
}

const baseAmmPacketData: object = {};

export const AmmPacketData = {
  encode(message: AmmPacketData, writer: Writer = Writer.create()): Writer {
    if (message.noData !== undefined) {
      NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
    }
    if (message.addLiquidityPacket !== undefined) {
      AddLiquidityPacketData.encode(
        message.addLiquidityPacket,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.createPoolPacket !== undefined) {
      CreatePoolPacketData.encode(
        message.createPoolPacket,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AmmPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAmmPacketData } as AmmPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.noData = NoData.decode(reader, reader.uint32());
          break;
        case 3:
          message.addLiquidityPacket = AddLiquidityPacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.createPoolPacket = CreatePoolPacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AmmPacketData {
    const message = { ...baseAmmPacketData } as AmmPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromJSON(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.addLiquidityPacket !== undefined &&
      object.addLiquidityPacket !== null
    ) {
      message.addLiquidityPacket = AddLiquidityPacketData.fromJSON(
        object.addLiquidityPacket
      );
    } else {
      message.addLiquidityPacket = undefined;
    }
    if (
      object.createPoolPacket !== undefined &&
      object.createPoolPacket !== null
    ) {
      message.createPoolPacket = CreatePoolPacketData.fromJSON(
        object.createPoolPacket
      );
    } else {
      message.createPoolPacket = undefined;
    }
    return message;
  },

  toJSON(message: AmmPacketData): unknown {
    const obj: any = {};
    message.noData !== undefined &&
      (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
    message.addLiquidityPacket !== undefined &&
      (obj.addLiquidityPacket = message.addLiquidityPacket
        ? AddLiquidityPacketData.toJSON(message.addLiquidityPacket)
        : undefined);
    message.createPoolPacket !== undefined &&
      (obj.createPoolPacket = message.createPoolPacket
        ? CreatePoolPacketData.toJSON(message.createPoolPacket)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<AmmPacketData>): AmmPacketData {
    const message = { ...baseAmmPacketData } as AmmPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromPartial(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.addLiquidityPacket !== undefined &&
      object.addLiquidityPacket !== null
    ) {
      message.addLiquidityPacket = AddLiquidityPacketData.fromPartial(
        object.addLiquidityPacket
      );
    } else {
      message.addLiquidityPacket = undefined;
    }
    if (
      object.createPoolPacket !== undefined &&
      object.createPoolPacket !== null
    ) {
      message.createPoolPacket = CreatePoolPacketData.fromPartial(
        object.createPoolPacket
      );
    } else {
      message.createPoolPacket = undefined;
    }
    return message;
  },
};

const baseNoData: object = {};

export const NoData = {
  encode(_: NoData, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NoData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNoData } as NoData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },

  toJSON(_: NoData): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<NoData>): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },
};

const baseCreatePoolPacketData: object = { poolParam: "", poolAssets: "" };

export const CreatePoolPacketData = {
  encode(
    message: CreatePoolPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolParam !== "") {
      writer.uint32(10).string(message.poolParam);
    }
    if (message.poolAssets !== "") {
      writer.uint32(18).string(message.poolAssets);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CreatePoolPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCreatePoolPacketData } as CreatePoolPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolParam = reader.string();
          break;
        case 2:
          message.poolAssets = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreatePoolPacketData {
    const message = { ...baseCreatePoolPacketData } as CreatePoolPacketData;
    if (object.poolParam !== undefined && object.poolParam !== null) {
      message.poolParam = String(object.poolParam);
    } else {
      message.poolParam = "";
    }
    if (object.poolAssets !== undefined && object.poolAssets !== null) {
      message.poolAssets = String(object.poolAssets);
    } else {
      message.poolAssets = "";
    }
    return message;
  },

  toJSON(message: CreatePoolPacketData): unknown {
    const obj: any = {};
    message.poolParam !== undefined && (obj.poolParam = message.poolParam);
    message.poolAssets !== undefined && (obj.poolAssets = message.poolAssets);
    return obj;
  },

  fromPartial(object: DeepPartial<CreatePoolPacketData>): CreatePoolPacketData {
    const message = { ...baseCreatePoolPacketData } as CreatePoolPacketData;
    if (object.poolParam !== undefined && object.poolParam !== null) {
      message.poolParam = object.poolParam;
    } else {
      message.poolParam = "";
    }
    if (object.poolAssets !== undefined && object.poolAssets !== null) {
      message.poolAssets = object.poolAssets;
    } else {
      message.poolAssets = "";
    }
    return message;
  },
};

const baseCreatePoolPacketAck: object = { poolId: 0 };

export const CreatePoolPacketAck = {
  encode(
    message: CreatePoolPacketAck,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolId !== 0) {
      writer.uint32(8).int32(message.poolId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CreatePoolPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCreatePoolPacketAck } as CreatePoolPacketAck;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolId = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreatePoolPacketAck {
    const message = { ...baseCreatePoolPacketAck } as CreatePoolPacketAck;
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = Number(object.poolId);
    } else {
      message.poolId = 0;
    }
    return message;
  },

  toJSON(message: CreatePoolPacketAck): unknown {
    const obj: any = {};
    message.poolId !== undefined && (obj.poolId = message.poolId);
    return obj;
  },

  fromPartial(object: DeepPartial<CreatePoolPacketAck>): CreatePoolPacketAck {
    const message = { ...baseCreatePoolPacketAck } as CreatePoolPacketAck;
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = object.poolId;
    } else {
      message.poolId = 0;
    }
    return message;
  },
};

const baseAddLiquidityPacketData: object = { amountVela: "", amountOther: "" };

export const AddLiquidityPacketData = {
  encode(
    message: AddLiquidityPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.amountVela !== "") {
      writer.uint32(10).string(message.amountVela);
    }
    if (message.amountOther !== "") {
      writer.uint32(18).string(message.amountOther);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AddLiquidityPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAddLiquidityPacketData } as AddLiquidityPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amountVela = reader.string();
          break;
        case 2:
          message.amountOther = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AddLiquidityPacketData {
    const message = { ...baseAddLiquidityPacketData } as AddLiquidityPacketData;
    if (object.amountVela !== undefined && object.amountVela !== null) {
      message.amountVela = String(object.amountVela);
    } else {
      message.amountVela = "";
    }
    if (object.amountOther !== undefined && object.amountOther !== null) {
      message.amountOther = String(object.amountOther);
    } else {
      message.amountOther = "";
    }
    return message;
  },

  toJSON(message: AddLiquidityPacketData): unknown {
    const obj: any = {};
    message.amountVela !== undefined && (obj.amountVela = message.amountVela);
    message.amountOther !== undefined &&
      (obj.amountOther = message.amountOther);
    return obj;
  },

  fromPartial(
    object: DeepPartial<AddLiquidityPacketData>
  ): AddLiquidityPacketData {
    const message = { ...baseAddLiquidityPacketData } as AddLiquidityPacketData;
    if (object.amountVela !== undefined && object.amountVela !== null) {
      message.amountVela = object.amountVela;
    } else {
      message.amountVela = "";
    }
    if (object.amountOther !== undefined && object.amountOther !== null) {
      message.amountOther = object.amountOther;
    } else {
      message.amountOther = "";
    }
    return message;
  },
};

const baseAddLiquidityPacketAck: object = { amountShares: "" };

export const AddLiquidityPacketAck = {
  encode(
    message: AddLiquidityPacketAck,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.amountShares !== "") {
      writer.uint32(10).string(message.amountShares);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AddLiquidityPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAddLiquidityPacketAck } as AddLiquidityPacketAck;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amountShares = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AddLiquidityPacketAck {
    const message = { ...baseAddLiquidityPacketAck } as AddLiquidityPacketAck;
    if (object.amountShares !== undefined && object.amountShares !== null) {
      message.amountShares = String(object.amountShares);
    } else {
      message.amountShares = "";
    }
    return message;
  },

  toJSON(message: AddLiquidityPacketAck): unknown {
    const obj: any = {};
    message.amountShares !== undefined &&
      (obj.amountShares = message.amountShares);
    return obj;
  },

  fromPartial(
    object: DeepPartial<AddLiquidityPacketAck>
  ): AddLiquidityPacketAck {
    const message = { ...baseAddLiquidityPacketAck } as AddLiquidityPacketAck;
    if (object.amountShares !== undefined && object.amountShares !== null) {
      message.amountShares = object.amountShares;
    } else {
      message.amountShares = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
