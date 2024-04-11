import argparse
import os
import shutil
import subprocess

PROTOS_PATH = "./proto"
XOR_GO_PATH = "./xor-go"
XOR_JAVA_PATH = "./xor-java"


def prepare_protos_for_xor_go(proto_folder_path: str, xor_go_path: str):
    proto_folders = [f for f in os.listdir(proto_folder_path) if os.path.isdir(os.path.join(proto_folder_path, f))]

    go_protos_path = f"{xor_go_path}/proto"
    if os.path.exists(go_protos_path):
        shutil.rmtree(go_protos_path)
    os.mkdir(go_protos_path)

    for proto_folder in proto_folders:
        src_path = os.path.join(proto_folder_path, proto_folder)
        dest_path = os.path.join(go_protos_path, proto_folder)

        if os.path.exists(dest_path):
            shutil.rmtree(dest_path)
        shutil.copytree(src_path, dest_path)

    command = (
        "protoc "
        "--go_out=. "
        "--go_opt=paths=source_relative "
        "--go-grpc_out=. "
        "--go-grpc_opt=paths=source_relative "
        "--go-grpc_opt=paths=source_relative "
        "xor-go/proto/*/*.proto"
    )
    subprocess.run(command, shell=True)


def prepare_protos_for_xor_java(proto_folder_path: str, xor_java_path: str):
    proto_folders = [f for f in os.listdir(proto_folder_path) if os.path.isdir(os.path.join(proto_folder_path, f))]

    for proto_folder in proto_folders:
        src_path = os.path.join(proto_folder_path, proto_folder)

        java_project_proto_path = f"{xor_java_path}/{proto_folder}-proto/src/main/proto"
        if os.path.exists(java_project_proto_path):
            shutil.rmtree(java_project_proto_path)
        shutil.copytree(src_path, java_project_proto_path)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--xor-java', type=bool, default=False, action=argparse.BooleanOptionalAction)
    parser.add_argument('--xor-go', type=bool, default=False, action=argparse.BooleanOptionalAction)
    args = parser.parse_args()

    if args.xor_java:
        prepare_protos_for_xor_java(PROTOS_PATH, XOR_JAVA_PATH)
    if args.xor_go:
        prepare_protos_for_xor_go(PROTOS_PATH, XOR_GO_PATH)


if __name__ == '__main__':
    main()
