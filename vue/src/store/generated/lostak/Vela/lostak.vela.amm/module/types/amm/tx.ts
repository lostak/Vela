/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "lostak.vela.amm";

export interface MsgSendCreatePool {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  poolParam: string;
  poolAssets: string;
}

export interface MsgSendCreatePoolResponse {}

export interface MsgSendAddLiquidity {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  amountVela: string;
  amountOther: string;
}

export interface MsgSendAddLiquidityResponse {}

export interface MsgSendRemoveLiquidity {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  amountShare: string;
  amountOther: string;
}

export interface MsgSendRemoveLiquidityResponse {}

const baseMsgSendCreatePool: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  poolParam: "",
  poolAssets: "",
};

export const MsgSendCreatePool = {
  encode(message: MsgSendCreatePool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.port !== "") {
      writer.uint32(18).string(message.port);
    }
    if (message.channelID !== "") {
      writer.uint32(26).string(message.channelID);
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(32).uint64(message.timeoutTimestamp);
    }
    if (message.poolParam !== "") {
      writer.uint32(42).string(message.poolParam);
    }
    if (message.poolAssets !== "") {
      writer.uint32(50).string(message.poolAssets);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendCreatePool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendCreatePool } as MsgSendCreatePool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.port = reader.string();
          break;
        case 3:
          message.channelID = reader.string();
          break;
        case 4:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.poolParam = reader.string();
          break;
        case 6:
          message.poolAssets = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendCreatePool {
    const message = { ...baseMsgSendCreatePool } as MsgSendCreatePool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = String(object.port);
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = String(object.channelID);
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = Number(object.timeoutTimestamp);
    } else {
      message.timeoutTimestamp = 0;
    }
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

  toJSON(message: MsgSendCreatePool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.poolParam !== undefined && (obj.poolParam = message.poolParam);
    message.poolAssets !== undefined && (obj.poolAssets = message.poolAssets);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSendCreatePool>): MsgSendCreatePool {
    const message = { ...baseMsgSendCreatePool } as MsgSendCreatePool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = object.port;
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = object.channelID;
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = object.timeoutTimestamp;
    } else {
      message.timeoutTimestamp = 0;
    }
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

const baseMsgSendCreatePoolResponse: object = {};

export const MsgSendCreatePoolResponse = {
  encode(
    _: MsgSendCreatePoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendCreatePoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendCreatePoolResponse,
    } as MsgSendCreatePoolResponse;
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

  fromJSON(_: any): MsgSendCreatePoolResponse {
    const message = {
      ...baseMsgSendCreatePoolResponse,
    } as MsgSendCreatePoolResponse;
    return message;
  },

  toJSON(_: MsgSendCreatePoolResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSendCreatePoolResponse>
  ): MsgSendCreatePoolResponse {
    const message = {
      ...baseMsgSendCreatePoolResponse,
    } as MsgSendCreatePoolResponse;
    return message;
  },
};

const baseMsgSendAddLiquidity: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  amountVela: "",
  amountOther: "",
};

export const MsgSendAddLiquidity = {
  encode(
    message: MsgSendAddLiquidity,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.port !== "") {
      writer.uint32(18).string(message.port);
    }
    if (message.channelID !== "") {
      writer.uint32(26).string(message.channelID);
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(32).uint64(message.timeoutTimestamp);
    }
    if (message.amountVela !== "") {
      writer.uint32(42).string(message.amountVela);
    }
    if (message.amountOther !== "") {
      writer.uint32(50).string(message.amountOther);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendAddLiquidity {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendAddLiquidity } as MsgSendAddLiquidity;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.port = reader.string();
          break;
        case 3:
          message.channelID = reader.string();
          break;
        case 4:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.amountVela = reader.string();
          break;
        case 6:
          message.amountOther = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendAddLiquidity {
    const message = { ...baseMsgSendAddLiquidity } as MsgSendAddLiquidity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = String(object.port);
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = String(object.channelID);
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = Number(object.timeoutTimestamp);
    } else {
      message.timeoutTimestamp = 0;
    }
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

  toJSON(message: MsgSendAddLiquidity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.amountVela !== undefined && (obj.amountVela = message.amountVela);
    message.amountOther !== undefined &&
      (obj.amountOther = message.amountOther);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSendAddLiquidity>): MsgSendAddLiquidity {
    const message = { ...baseMsgSendAddLiquidity } as MsgSendAddLiquidity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = object.port;
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = object.channelID;
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = object.timeoutTimestamp;
    } else {
      message.timeoutTimestamp = 0;
    }
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

const baseMsgSendAddLiquidityResponse: object = {};

export const MsgSendAddLiquidityResponse = {
  encode(
    _: MsgSendAddLiquidityResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendAddLiquidityResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendAddLiquidityResponse,
    } as MsgSendAddLiquidityResponse;
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

  fromJSON(_: any): MsgSendAddLiquidityResponse {
    const message = {
      ...baseMsgSendAddLiquidityResponse,
    } as MsgSendAddLiquidityResponse;
    return message;
  },

  toJSON(_: MsgSendAddLiquidityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSendAddLiquidityResponse>
  ): MsgSendAddLiquidityResponse {
    const message = {
      ...baseMsgSendAddLiquidityResponse,
    } as MsgSendAddLiquidityResponse;
    return message;
  },
};

const baseMsgSendRemoveLiquidity: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  amountShare: "",
  amountOther: "",
};

export const MsgSendRemoveLiquidity = {
  encode(
    message: MsgSendRemoveLiquidity,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.port !== "") {
      writer.uint32(18).string(message.port);
    }
    if (message.channelID !== "") {
      writer.uint32(26).string(message.channelID);
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(32).uint64(message.timeoutTimestamp);
    }
    if (message.amountShare !== "") {
      writer.uint32(42).string(message.amountShare);
    }
    if (message.amountOther !== "") {
      writer.uint32(50).string(message.amountOther);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendRemoveLiquidity {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendRemoveLiquidity } as MsgSendRemoveLiquidity;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.port = reader.string();
          break;
        case 3:
          message.channelID = reader.string();
          break;
        case 4:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.amountShare = reader.string();
          break;
        case 6:
          message.amountOther = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendRemoveLiquidity {
    const message = { ...baseMsgSendRemoveLiquidity } as MsgSendRemoveLiquidity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = String(object.port);
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = String(object.channelID);
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = Number(object.timeoutTimestamp);
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.amountShare !== undefined && object.amountShare !== null) {
      message.amountShare = String(object.amountShare);
    } else {
      message.amountShare = "";
    }
    if (object.amountOther !== undefined && object.amountOther !== null) {
      message.amountOther = String(object.amountOther);
    } else {
      message.amountOther = "";
    }
    return message;
  },

  toJSON(message: MsgSendRemoveLiquidity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.amountShare !== undefined &&
      (obj.amountShare = message.amountShare);
    message.amountOther !== undefined &&
      (obj.amountOther = message.amountOther);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendRemoveLiquidity>
  ): MsgSendRemoveLiquidity {
    const message = { ...baseMsgSendRemoveLiquidity } as MsgSendRemoveLiquidity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = object.port;
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = object.channelID;
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = object.timeoutTimestamp;
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.amountShare !== undefined && object.amountShare !== null) {
      message.amountShare = object.amountShare;
    } else {
      message.amountShare = "";
    }
    if (object.amountOther !== undefined && object.amountOther !== null) {
      message.amountOther = object.amountOther;
    } else {
      message.amountOther = "";
    }
    return message;
  },
};

const baseMsgSendRemoveLiquidityResponse: object = {};

export const MsgSendRemoveLiquidityResponse = {
  encode(
    _: MsgSendRemoveLiquidityResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendRemoveLiquidityResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendRemoveLiquidityResponse,
    } as MsgSendRemoveLiquidityResponse;
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

  fromJSON(_: any): MsgSendRemoveLiquidityResponse {
    const message = {
      ...baseMsgSendRemoveLiquidityResponse,
    } as MsgSendRemoveLiquidityResponse;
    return message;
  },

  toJSON(_: MsgSendRemoveLiquidityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSendRemoveLiquidityResponse>
  ): MsgSendRemoveLiquidityResponse {
    const message = {
      ...baseMsgSendRemoveLiquidityResponse,
    } as MsgSendRemoveLiquidityResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  SendCreatePool(
    request: MsgSendCreatePool
  ): Promise<MsgSendCreatePoolResponse>;
  SendAddLiquidity(
    request: MsgSendAddLiquidity
  ): Promise<MsgSendAddLiquidityResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SendRemoveLiquidity(
    request: MsgSendRemoveLiquidity
  ): Promise<MsgSendRemoveLiquidityResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SendCreatePool(
    request: MsgSendCreatePool
  ): Promise<MsgSendCreatePoolResponse> {
    const data = MsgSendCreatePool.encode(request).finish();
    const promise = this.rpc.request(
      "lostak.vela.amm.Msg",
      "SendCreatePool",
      data
    );
    return promise.then((data) =>
      MsgSendCreatePoolResponse.decode(new Reader(data))
    );
  }

  SendAddLiquidity(
    request: MsgSendAddLiquidity
  ): Promise<MsgSendAddLiquidityResponse> {
    const data = MsgSendAddLiquidity.encode(request).finish();
    const promise = this.rpc.request(
      "lostak.vela.amm.Msg",
      "SendAddLiquidity",
      data
    );
    return promise.then((data) =>
      MsgSendAddLiquidityResponse.decode(new Reader(data))
    );
  }

  SendRemoveLiquidity(
    request: MsgSendRemoveLiquidity
  ): Promise<MsgSendRemoveLiquidityResponse> {
    const data = MsgSendRemoveLiquidity.encode(request).finish();
    const promise = this.rpc.request(
      "lostak.vela.amm.Msg",
      "SendRemoveLiquidity",
      data
    );
    return promise.then((data) =>
      MsgSendRemoveLiquidityResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
