package ride

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wavesplatform/gowaves/pkg/ride/meta"
)

func fl(functions ...meta.Function) []meta.Function {
	return append(make([]meta.Function, 0), functions...)
}

func f(name string, types ...meta.Type) meta.Function {
	return meta.Function{
		Name:      name,
		Arguments: append(make([]meta.Type, 0, len(types)), types...),
	}
}

func u(types ...meta.SimpleType) meta.UnionType {
	return append(make([]meta.SimpleType, 0), types...)
}
func l(t meta.Type) meta.ListType {
	return meta.ListType{Inner: t}
}

/*
# Test 0

{-# STDLIB_VERSION 3 #-}
{-# CONTENT_TYPE DAPP #-}
{-# SCRIPT_TYPE ACCOUNT #-}

@Callable(i)
func call1(a: ByteVector|String, b: Int|Boolean) = ScriptResult(WriteSet(nil), TransferSet(nil))

@Callable(i)
func call2(a: Int|Boolean|ByteVector|String) = ScriptResult(WriteSet(nil), TransferSet(nil))

@Verifier(tx)
func verify() = true
*/

/*
# Test 1

{-# STDLIB_VERSION 5 #-}
{-# CONTENT_TYPE DAPP #-}
{-# SCRIPT_TYPE ACCOUNT #-}

@Callable(i)
func call1(n: Int, b: Boolean, s: String, a: ByteVector) = []


@Callable(i)
func call2(li: List[Int], lsb: List[ByteVector|String], sl: List[String]) = []

@Verifier(tx)
func verify() = true
*/
func TestConvertMetaV1(t *testing.T) {
	for _, test := range []struct {
		comment string
		source  string
		v       int
		f       []meta.Function
	}{
		{`V3: no meta`, "AAIDAAAAAAAAAAAAAAABAQAAABFnZXRQcmV2aW91c0Fuc3dlcgAAAAEAAAAHYWRkcmVzcwUAAAAHYWRkcmVzcwAAAAIAAAABaQEAAAAGdGVsbG1lAAAAAQAAAAhxdWVzdGlvbgQAAAAGYW5zd2VyCQEAAAARZ2V0UHJldmlvdXNBbnN3ZXIAAAABBQAAAAhxdWVzdGlvbgkBAAAACFdyaXRlU2V0AAAAAQkABEwAAAACCQEAAAAJRGF0YUVudHJ5AAAAAgkAASwAAAACBQAAAAZhbnN3ZXICAAAAAl9xBQAAAAhxdWVzdGlvbgkABEwAAAACCQEAAAAJRGF0YUVudHJ5AAAAAgkAASwAAAACBQAAAAZhbnN3ZXICAAAAAl9hBQAAAAZhbnN3ZXIFAAAAA25pbAAAAAppbnZvY2F0aW9uAQAAAAdkZWZhdWx0AAAAAAQAAAAHc2VuZGVyMAgIBQAAAAppbnZvY2F0aW9uAAAABmNhbGxlcgAAAAVieXRlcwkBAAAACFdyaXRlU2V0AAAAAQkABEwAAAACCQEAAAAJRGF0YUVudHJ5AAAAAgIAAAABYQIAAAABYgkABEwAAAACCQEAAAAJRGF0YUVudHJ5AAAAAgIAAAAGc2VuZGVyBQAAAAdzZW5kZXIwBQAAAANuaWwAAAABAAAAAnR4AQAAAAZ2ZXJpZnkAAAAACQAAAAAAAAIJAQAAABFnZXRQcmV2aW91c0Fuc3dlcgAAAAEJAAQlAAAAAQgFAAAAAnR4AAAABnNlbmRlcgIAAAABMcP91gY=",
			0, nil},
		{
			`V3: wallet example`, "AAIDAAAAAAAAAAkIARIAEgMKAQEAAAAAAAAAAgAAAAFpAQAAAAdkZXBvc2l0AAAAAAQAAAADcG10CQEAAAAHZXh0cmFjdAAAAAEIBQAAAAFpAAAAB3BheW1lbnQDCQEAAAAJaXNEZWZpbmVkAAAAAQgFAAAAA3BtdAAAAAdhc3NldElkCQAAAgAAAAECAAAAIWNhbiBob2RsIHdhdmVzIG9ubHkgYXQgdGhlIG1vbWVudAQAAAAKY3VycmVudEtleQkAAlgAAAABCAgFAAAAAWkAAAAGY2FsbGVyAAAABWJ5dGVzBAAAAA1jdXJyZW50QW1vdW50BAAAAAckbWF0Y2gwCQAEGgAAAAIFAAAABHRoaXMFAAAACmN1cnJlbnRLZXkDCQAAAQAAAAIFAAAAByRtYXRjaDACAAAAA0ludAQAAAABYQUAAAAHJG1hdGNoMAUAAAABYQAAAAAAAAAAAAQAAAAJbmV3QW1vdW50CQAAZAAAAAIFAAAADWN1cnJlbnRBbW91bnQIBQAAAANwbXQAAAAGYW1vdW50CQEAAAAIV3JpdGVTZXQAAAABCQAETAAAAAIJAQAAAAlEYXRhRW50cnkAAAACBQAAAApjdXJyZW50S2V5BQAAAAluZXdBbW91bnQFAAAAA25pbAAAAAFpAQAAAAh3aXRoZHJhdwAAAAEAAAAGYW1vdW50BAAAAApjdXJyZW50S2V5CQACWAAAAAEICAUAAAABaQAAAAZjYWxsZXIAAAAFYnl0ZXMEAAAADWN1cnJlbnRBbW91bnQEAAAAByRtYXRjaDAJAAQaAAAAAgUAAAAEdGhpcwUAAAAKY3VycmVudEtleQMJAAABAAAAAgUAAAAHJG1hdGNoMAIAAAADSW50BAAAAAFhBQAAAAckbWF0Y2gwBQAAAAFhAAAAAAAAAAAABAAAAAluZXdBbW91bnQJAABlAAAAAgUAAAANY3VycmVudEFtb3VudAUAAAAGYW1vdW50AwkAAGYAAAACAAAAAAAAAAAABQAAAAZhbW91bnQJAAACAAAAAQIAAAAeQ2FuJ3Qgd2l0aGRyYXcgbmVnYXRpdmUgYW1vdW50AwkAAGYAAAACAAAAAAAAAAAABQAAAAluZXdBbW91bnQJAAACAAAAAQIAAAASTm90IGVub3VnaCBiYWxhbmNlCQEAAAAMU2NyaXB0UmVzdWx0AAAAAgkBAAAACFdyaXRlU2V0AAAAAQkABEwAAAACCQEAAAAJRGF0YUVudHJ5AAAAAgUAAAAKY3VycmVudEtleQUAAAAJbmV3QW1vdW50BQAAAANuaWwJAQAAAAtUcmFuc2ZlclNldAAAAAEJAARMAAAAAgkBAAAADlNjcmlwdFRyYW5zZmVyAAAAAwgFAAAAAWkAAAAGY2FsbGVyBQAAAAZhbW91bnQFAAAABHVuaXQFAAAAA25pbAAAAAEAAAACdHgBAAAABnZlcmlmeQAAAAAH+vmYeA==",
			1, fl(f("deposit"), f("withdraw", meta.Int)),
		},
		{
			`V3: 8ball example`, "AAIDAAAAAAAAAAcIARIDCgEIAAAABAAAAAAMYW5zd2Vyc0NvdW50AAAAAAAAAAAUAAAAAAdhbnN3ZXJzCQAETAAAAAICAAAADkl0IGlzIGNlcnRhaW4uCQAETAAAAAICAAAAE0l0IGlzIGRlY2lkZWRseSBzby4JAARMAAAAAgIAAAAQV2l0aG91dCBhIGRvdWJ0LgkABEwAAAACAgAAABFZZXMgLSBkZWZpbml0ZWx5LgkABEwAAAACAgAAABNZb3UgbWF5IHJlbHkgb24gaXQuCQAETAAAAAICAAAAEUFzIEkgc2VlIGl0LCB5ZXMuCQAETAAAAAICAAAADE1vc3QgbGlrZWx5LgkABEwAAAACAgAAAA1PdXRsb29rIGdvb2QuCQAETAAAAAICAAAABFllcy4JAARMAAAAAgIAAAATU2lnbnMgcG9pbnQgdG8geWVzLgkABEwAAAACAgAAABZSZXBseSBoYXp5LCB0cnkgYWdhaW4uCQAETAAAAAICAAAAEEFzayBhZ2FpbiBsYXRlci4JAARMAAAAAgIAAAAYQmV0dGVyIG5vdCB0ZWxsIHlvdSBub3cuCQAETAAAAAICAAAAE0Nhbm5vdCBwcmVkaWN0IG5vdy4JAARMAAAAAgIAAAAaQ29uY2VudHJhdGUgYW5kIGFzayBhZ2Fpbi4JAARMAAAAAgIAAAASRG9uJ3QgY291bnQgb24gaXQuCQAETAAAAAICAAAAD015IHJlcGx5IGlzIG5vLgkABEwAAAACAgAAABJNeSBzb3VyY2VzIHNheSBuby4JAARMAAAAAgIAAAAUT3V0bG9vayBub3Qgc28gZ29vZC4JAARMAAAAAgIAAAAOVmVyeSBkb3VidGZ1bC4FAAAAA25pbAEAAAAJZ2V0QW5zd2VyAAAAAgAAAAhxdWVzdGlvbgAAAA5wcmV2aW91c0Fuc3dlcgQAAAAEaGFzaAkAAfcAAAABCQABmwAAAAEJAAEsAAAAAgUAAAAIcXVlc3Rpb24FAAAADnByZXZpb3VzQW5zd2VyBAAAAAVpbmRleAkABLEAAAABBQAAAARoYXNoCQABkQAAAAIFAAAAB2Fuc3dlcnMJAABqAAAAAgUAAAAFaW5kZXgFAAAADGFuc3dlcnNDb3VudAEAAAARZ2V0UHJldmlvdXNBbnN3ZXIAAAABAAAAB2FkZHJlc3MEAAAAByRtYXRjaDAJAAQdAAAAAgUAAAAEdGhpcwkAASwAAAACBQAAAAdhZGRyZXNzAgAAAAJfYQMJAAABAAAAAgUAAAAHJG1hdGNoMAIAAAAGU3RyaW5nBAAAAAFhBQAAAAckbWF0Y2gwBQAAAAFhBQAAAAdhZGRyZXNzAAAAAQAAAAFpAQAAAAZ0ZWxsbWUAAAABAAAACHF1ZXN0aW9uBAAAAA1jYWxsZXJBZGRyZXNzCQACWAAAAAEICAUAAAABaQAAAAZjYWxsZXIAAAAFYnl0ZXMEAAAABmFuc3dlcgkBAAAACWdldEFuc3dlcgAAAAIFAAAACHF1ZXN0aW9uCQEAAAARZ2V0UHJldmlvdXNBbnN3ZXIAAAABBQAAAA1jYWxsZXJBZGRyZXNzCQEAAAAIV3JpdGVTZXQAAAABCQAETAAAAAIJAQAAAAlEYXRhRW50cnkAAAACCQABLAAAAAIFAAAADWNhbGxlckFkZHJlc3MCAAAAAl9xBQAAAAhxdWVzdGlvbgkABEwAAAACCQEAAAAJRGF0YUVudHJ5AAAAAgkAASwAAAACBQAAAA1jYWxsZXJBZGRyZXNzAgAAAAJfYQUAAAAGYW5zd2VyBQAAAANuaWwAAAAAOvBJMA==",
			1, fl(f("tellme", meta.String)),
		},
		{
			`V3: casino example`, "AAIDAAAAAAAAAA4IARIFCgMIAQESAwoBCAAAAAQAAAAABm9yYWNsZQkBAAAAB2V4dHJhY3QAAAABCQEAAAARYWRkcmVzc0Zyb21TdHJpbmcAAAABAgAAAA8kT1JBQ0xFX0FERFJFU1MAAAAABm1pbkJldAAAAAAAAvrwgAAAAAAJbWF4U3VtQmV0AAAAAAA7msoAAQAAAA1jYWxjV2luQW1vdW50AAAAAwAAAAR0aGlzAAAAA2tleQAAAAVrb2VmZgQAAAAHJG1hdGNoMAkABBoAAAACBQAAAAR0aGlzBQAAAANrZXkDCQAAAQAAAAIFAAAAByRtYXRjaDACAAAAA0ludAQAAAABYQUAAAAHJG1hdGNoMAkAAGgAAAACBQAAAAFhBQAAAAVrb2VmZgAAAAAAAAAAAAAAAAIAAAABaQEAAAADYmV0AAAAAwAAAAVyb3VuZAAAAAlndWVzc1R5cGUAAAAKZ3Vlc3NWYWx1ZQQAAAADcG10CQEAAAAHZXh0cmFjdAAAAAEIBQAAAAFpAAAAB3BheW1lbnQDCQEAAAAJaXNEZWZpbmVkAAAAAQgFAAAAA3BtdAAAAAdhc3NldElkCQAAAgAAAAECAAAAHEJldHMgb25seSBpbiBXYXZlcyBzdXBwb3J0ZWQDCQAAZgAAAAIFAAAABm1pbkJldAgFAAAAA3BtdAAAAAZhbW91bnQJAAACAAAAAQkAASwAAAACAgAAAClZb3VyIEJldCBhbW91bnQgaXMgbGVzcyB0aGVuIG1pbmltYWwgYmV0IAkAAaQAAAABBQAAAAZtaW5CZXQDAwkBAAAACWlzRGVmaW5lZAAAAAEJAAQbAAAAAgUAAAAGb3JhY2xlCQABLAAAAAIFAAAABXJvdW5kAgAAAAVfc3RvcAYJAQAAAAlpc0RlZmluZWQAAAABCQAEHQAAAAIFAAAABm9yYWNsZQUAAAAFcm91bmQJAAACAAAAAQIAAAAcVGhpcyByb3VuZCBpcyBhbHJlYWR5IHBsYXllZAQAAAAMcm91bmRCZXRzS2V5CQABLAAAAAIFAAAABXJvdW5kAgAAAAhfc3VtQmV0cwQAAAAKY3VyU3VtQmV0cwQAAAAHJG1hdGNoMAkABBoAAAACBQAAAAR0aGlzBQAAAAxyb3VuZEJldHNLZXkDCQAAAQAAAAIFAAAAByRtYXRjaDACAAAAA0ludAQAAAABYQUAAAAHJG1hdGNoMAUAAAABYQAAAAAAAAAAAAQAAAAKbmV3U3VtQmV0cwkAAGQAAAACBQAAAApjdXJTdW1CZXRzCAUAAAADcG10AAAABmFtb3VudAMJAABmAAAAAgUAAAAKbmV3U3VtQmV0cwUAAAAJbWF4U3VtQmV0CQAAAgAAAAEJAAEsAAAAAgkAASwAAAACCQABLAAAAAICAAAAIU1heGltdW0gYW1vdW50IG9mIGJldHMgZm9yIHJvdW5kIAkAAaQAAAABBQAAAAltYXhTdW1CZXQCAAAAFS4gV2l0aCB5b3VyIGJldCBpdCdzIAkAAaQAAAABBQAAAApuZXdTdW1CZXRzBAAAAAZiZXRLZXkJAAEsAAAAAgkAASwAAAACCQABLAAAAAIJAAEsAAAAAgkAASwAAAACCQABLAAAAAIJAAJYAAAAAQgIBQAAAAFpAAAABmNhbGxlcgAAAAVieXRlcwIAAAABXwUAAAAFcm91bmQCAAAAAV8JAAGkAAAAAQUAAAAJZ3Vlc3NUeXBlAgAAAAFfCQABpAAAAAEFAAAACmd1ZXNzVmFsdWUEAAAADGN1ckJldEFtb3VudAQAAAAHJG1hdGNoMAkABBoAAAACBQAAAAR0aGlzBQAAAAZiZXRLZXkDCQAAAQAAAAIFAAAAByRtYXRjaDACAAAAA0ludAQAAAABYQUAAAAHJG1hdGNoMAUAAAABYQAAAAAAAAAAAAQAAAAMbmV3QmV0QW1vdW50CQAAZAAAAAIFAAAADGN1ckJldEFtb3VudAgFAAAAA3BtdAAAAAZhbW91bnQJAQAAAAhXcml0ZVNldAAAAAEJAARMAAAAAgkBAAAACURhdGFFbnRyeQAAAAIFAAAABmJldEtleQUAAAAMbmV3QmV0QW1vdW50CQAETAAAAAIJAQAAAAlEYXRhRW50cnkAAAACBQAAAAxyb3VuZEJldHNLZXkFAAAACm5ld1N1bUJldHMFAAAAA25pbAAAAAFpAQAAAAh3aXRoZHJhdwAAAAEAAAAFcm91bmQEAAAACmJldEtleVBhcnQJAAEsAAAAAgkAASwAAAACCQACWAAAAAEICAUAAAABaQAAAAZjYWxsZXIAAAAFYnl0ZXMCAAAAAV8FAAAABXJvdW5kBAAAAAt3aXRoZHJhd0tleQkAASwAAAACBQAAAApiZXRLZXlQYXJ0AgAAAAlfd2l0aGRyYXcDCQEAAAAJaXNEZWZpbmVkAAAAAQkABBoAAAACBQAAAAR0aGlzBQAAAAt3aXRoZHJhd0tleQkAAAIAAAABAgAAAB9Zb3UgaGF2ZSBhbHJlYWR5IGdvdCB5b3VyIHByaXplBAAAAAp2YWxDb21wbGV4CQEAAAARQGV4dHJOYXRpdmUoMTA1MykAAAACBQAAAAZvcmFjbGUFAAAABXJvdW5kBAAAAAZ3aW5OdW0JAQAAAA1wYXJzZUludFZhbHVlAAAAAQkAATAAAAACCQABLwAAAAIFAAAACnZhbENvbXBsZXgAAAAAAAAAAAIAAAAAAAAAAAAEAAAAC3dpblJlZEJsYWNrCQEAAAANcGFyc2VJbnRWYWx1ZQAAAAEJAAEwAAAAAgkAAS8AAAACBQAAAAp2YWxDb21wbGV4AAAAAAAAAAADAAAAAAAAAAACBAAAAAp3aW5FdmVuT2RkCQEAAAANcGFyc2VJbnRWYWx1ZQAAAAEJAAEwAAAAAgkAAS8AAAACBQAAAAp2YWxDb21wbGV4AAAAAAAAAAAEAAAAAAAAAAADBAAAAAt3aW5EZXNrSGFsZgkBAAAADXBhcnNlSW50VmFsdWUAAAABCQABMAAAAAIJAAEvAAAAAgUAAAAKdmFsQ29tcGxleAAAAAAAAAAABQAAAAAAAAAABAQAAAAMd2luRGVza1RoaXJkCQEAAAANcGFyc2VJbnRWYWx1ZQAAAAEJAAEwAAAAAgkAAS8AAAACBQAAAAp2YWxDb21wbGV4AAAAAAAAAAAGAAAAAAAAAAAFBAAAAAZ3aW5Sb3cJAQAAAA1wYXJzZUludFZhbHVlAAAAAQkAATAAAAACCQABLwAAAAIFAAAACnZhbENvbXBsZXgAAAAAAAAAAAcAAAAAAAAAAAYEAAAACXdpbkFtb3VudAkAAGQAAAACCQAAZAAAAAIJAABkAAAAAgkAAGQAAAACCQAAZAAAAAIJAQAAAA1jYWxjV2luQW1vdW50AAAAAwUAAAAEdGhpcwkAASwAAAACCQABLAAAAAIFAAAACmJldEtleVBhcnQCAAAAA18wXwkAAaQAAAABBQAAAAZ3aW5OdW0AAAAAAAAAACQJAQAAAA1jYWxjV2luQW1vdW50AAAAAwUAAAAEdGhpcwkAASwAAAACCQABLAAAAAIFAAAACmJldEtleVBhcnQCAAAAA18xXwkAAaQAAAABBQAAAAt3aW5SZWRCbGFjawAAAAAAAAAAAgkBAAAADWNhbGNXaW5BbW91bnQAAAADBQAAAAR0aGlzCQABLAAAAAIJAAEsAAAAAgUAAAAKYmV0S2V5UGFydAIAAAADXzJfCQABpAAAAAEFAAAACndpbkV2ZW5PZGQAAAAAAAAAAAIJAQAAAA1jYWxjV2luQW1vdW50AAAAAwUAAAAEdGhpcwkAASwAAAACCQABLAAAAAIFAAAACmJldEtleVBhcnQCAAAAA18zXwkAAaQAAAABBQAAAAt3aW5EZXNrSGFsZgAAAAAAAAAAAgkBAAAADWNhbGNXaW5BbW91bnQAAAADBQAAAAR0aGlzCQABLAAAAAIJAAEsAAAAAgUAAAAKYmV0S2V5UGFydAIAAAADXzRfCQABpAAAAAEFAAAADHdpbkRlc2tUaGlyZAAAAAAAAAAAAwkBAAAADWNhbGNXaW5BbW91bnQAAAADBQAAAAR0aGlzCQABLAAAAAIJAAEsAAAAAgUAAAAKYmV0S2V5UGFydAIAAAADXzVfCQABpAAAAAEFAAAABndpblJvdwAAAAAAAAAAAwMJAAAAAAAAAgUAAAAJd2luQW1vdW50AAAAAAAAAAAACQAAAgAAAAECAAAAGllvdSB3b24gbm90aGluZyB0aGlzIHJvdW5kCQEAAAAMU2NyaXB0UmVzdWx0AAAAAgkBAAAACFdyaXRlU2V0AAAAAQkABEwAAAACCQEAAAAJRGF0YUVudHJ5AAAAAgUAAAALd2l0aGRyYXdLZXkFAAAACXdpbkFtb3VudAUAAAADbmlsCQEAAAALVHJhbnNmZXJTZXQAAAABCQAETAAAAAIJAQAAAA5TY3JpcHRUcmFuc2ZlcgAAAAMIBQAAAAFpAAAABmNhbGxlcgUAAAAJd2luQW1vdW50BQAAAAR1bml0BQAAAANuaWwAAAABAAAAAnR4AQAAAAZ2ZXJpZnkAAAAACQAB9AAAAAMIBQAAAAJ0eAAAAAlib2R5Qnl0ZXMJAAGRAAAAAggFAAAAAnR4AAAABnByb29mcwAAAAAAAAAAAAEAAAAg2AGRxriAuyMT3c9AcH6CQ43JiN5WSLBSu4r1qoLU0WKqOYr1",
			1, fl(f("bet", meta.String, meta.Int, meta.Int), f("withdraw", meta.String)),
		},
		{
			`V3: test 0`, "AAIDAAAAAAAAAA0IARIECgIKBRIDCgEPAAAAAAAAAAIAAAABaQEAAAAFY2FsbDEAAAACAAAAAWEAAAABYgkBAAAADFNjcmlwdFJlc3VsdAAAAAIJAQAAAAhXcml0ZVNldAAAAAEFAAAAA25pbAkBAAAAC1RyYW5zZmVyU2V0AAAAAQUAAAADbmlsAAAAAWkBAAAABWNhbGwyAAAAAQAAAAFhCQEAAAAMU2NyaXB0UmVzdWx0AAAAAgkBAAAACFdyaXRlU2V0AAAAAQUAAAADbmlsCQEAAAALVHJhbnNmZXJTZXQAAAABBQAAAANuaWwAAAABAAAAAnR4AQAAAAZ2ZXJpZnkAAAAABg8tLpI=",
			1, fl(f("call1", u(meta.Bytes, meta.String), u(meta.Int, meta.Boolean)), f("call2", u(meta.Int, meta.Bytes, meta.Boolean, meta.String))),
		},
		{
			`V5: test 1`, "AAIFAAAAAAAAABEIAhIGCgQBBAgCEgUKAxEaGAAAAAAAAAACAAAAAWkBAAAABWNhbGwxAAAABAAAAAFuAAAAAWIAAAABcwAAAAFhBQAAAANuaWwAAAABaQEAAAAFY2FsbDIAAAADAAAAAmxpAAAAA2xzYgAAAAJzbAUAAAADbmlsAAAAAQAAAAJ0eAEAAAAGdmVyaWZ5AAAAAAaVNfq8",
			2, fl(f("call1", meta.Int, meta.Boolean, meta.String, meta.Bytes), f("call2", l(meta.Int), l(u(meta.Bytes, meta.String)), l(meta.String))),
		},
	} {
		src, err := base64.StdEncoding.DecodeString(test.source)
		require.NoError(t, err, test.comment)

		tree, err := Parse(src)
		require.NoError(t, err, test.comment)
		assert.NotNil(t, tree.Meta, test.comment)

		assert.Equal(t, test.v, tree.Meta.Version, test.comment)
		assert.Equal(t, test.f, tree.Meta.Functions, test.comment)
	}
}
