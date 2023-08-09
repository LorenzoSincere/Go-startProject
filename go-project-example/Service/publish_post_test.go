package Service

import (
	"os"
	"startProject/go-project-example/Repository"
	"startProject/go-project-example/Util"
	"testing"
)

func TestMain(m *testing.M) {
	if err := Repository.Init(); err != nil {
		os.Exit(1)
	}
	if err := Util.InitLogger(); err != nil {
		os.Exit(1)
	}
	m.Run()
}

func TestPublishPost(t *testing.T) {
	type args struct {
		topicId int64
		userId  int64
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试发布回帖",
			args: args{
				topicId: 1,
				userId:  2,
				content: "再次回帖",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := PublishPost(tt.args.topicId, tt.args.userId, tt.args.content)

			if (err != nil) != tt.wantErr {
				t.Errorf("publishPost() error: %v, wantError: %v", err, tt.wantErr)
				return
			}
		})
	}
}

func BenchmarkPublishPost(b *testing.B) {
	type args struct {
		topicId int64
		userId  int64
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试发布回帖",
			args: args{
				topicId: 1,
				userId:  2,
				content: "再次回帖",
			},
			wantErr: false,
		},
	}

	b.ResetTimer()
	for _, tt := range tests {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, err := PublishPost(tt.args.topicId, tt.args.userId, tt.args.content)
				if (err != nil) != tt.wantErr {
					b.Errorf("publishPost() error: %v, wantError: %v", err, tt.wantErr)
					return
				}
			}

		})
	}
}
