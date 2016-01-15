package main

type ForceTransformer interface {
	Transform(data []byte) (forces []Force, err error)
}
