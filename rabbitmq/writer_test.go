package rabbitmq

import (
	"reflect"
	"testing"

	"github.com/streadway/amqp"
)

func TestNewWriter(t *testing.T) {
	type args struct {
		conf  Config
		queue string
	}
	tests := []struct {
		name    string
		args    args
		want    *Writer
		wantErr bool
	}{
		struct {
			name    string
			args    args
			want    *Writer
			wantErr bool
		}{name: "Test Failed initial", args: args{
			queue: "some queue",
			conf: Config{
				Port:     5672,
				Address:  "localhost",
				Password: "some_password",
				User:     "root",
			},
		}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWriter(tt.args.conf, tt.args.queue)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWriter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriter_Write(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		w       *Writer
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
		struct {
			name    string
			w       *Writer
			args    args
			wantN   int
			wantErr bool
		}{
			name: "Failed Publish data",
			w: &Writer{
				Conn:  &amqp.Connection{},
				Queue: amqp.Queue{},
			},
			args: args{
				data: []byte("29 23 23"),
			},
			wantN: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := tt.w.Write(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Writer.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Writer.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestWriter_Close(t *testing.T) {
	tests := []struct {
		name    string
		w       *Writer
		wantErr bool
	}{
		// TODO: Add test cases.
		struct {
			name    string
			w       *Writer
			wantErr bool
		}{name: "Close the mq", w: &Writer{
			Conn:  &amqp.Connection{},
			Queue: amqp.Queue{},
		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Writer.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
