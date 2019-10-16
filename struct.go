package main

// Storage 는
type Storage struct {
	ID      string // 스토리지 ID
	Windows string // 스토리지의 Windows 물리적 경로
	Linux   string // 스토리지의 Linux 물리적 경로
	MacOS   string // 스토리지의 macOS 물리적 경로
}

// Attribute 는 파일 타입과 상관없이 수 많은 자료를 저장하기 위해서 사용한다.
type Attribute struct {
	Key   string // 속성 키
	Value string // 속성 값
}

// Item 은 라이브러리의 에셋 자료구조이다.
type Item struct {
	ID          string      // ID
	Author      string      // 에셋을 제작한 사람
	Tags        []string    // 태그리스트
	Description string      // 에셋에 대한 추가 정보. 에셋의 제약, 사용전 알아야 할 특징
	Thumbimg    string      // 썸네일 이미지 주소
	Thumbmov    string      // 썸네일 영상 주소
	Inputpath   string      // 최초 등록되는 경로
	Outputpath  string      // 저장되는 경로
	Type        string      // maya, source, houdini, blender, nuke ..  같은 형태인가.
	Status      string      // 상태(에러, done, wip)
	Log         string      // 데이터를 처리할 때 생성되는 로그
	CreateTime  string      // Item 생성 시간
	Updatetime  string      // UTC 타임으로 들어가도록 하기.
	UsingRate   int64       // 사용 빈도 수
	Storage                 // Item이 저장되는 스토리지 정보
	Attributes  []Attribute // 해상도, 속성, 메타데이터 등의 파일정보
}