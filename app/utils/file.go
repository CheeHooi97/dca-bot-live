package utils

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type File struct {
	r2Client *s3.Client
}

// func NewFile() (*File, error) {
// 	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
// 		awsConfig.WithRegion(config.R2Region),
// 		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.R2AccessKey, config.R2AccessSecret, "")),
// 	)
// 	if err != nil {
// 		log.Fatalf("failed to load config for R2: %v", err)
// 	}

// 	r2Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
// 		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", config.R2AccountId))
// 		o.UsePathStyle = true
// 	})

// 	return &File{
// 		r2Client: r2Client,
// 	}, nil
// }

// func (f *File) UploadFile(ctx context.Context, userId, homeId, name string, num int, fileHeader *multipart.FileHeader, mediaBytes []byte) (string, string, error) {
// 	mediaType := mimetype.Detect(mediaBytes).String()
// 	filename := "unknown"
// 	if fileHeader != nil {
// 		filename = filepath.Base(fileHeader.Filename)
// 	}

// 	var key string
// 	var file string

// 	parts := strings.Split(mediaType, "/")
// 	if len(parts) == 2 {
// 		file = fmt.Sprintf(".%s", parts[1])
// 	}

// 	if userId != "" {
// 		key = fmt.Sprintf("%s/profile_%s%s", config.Env, userId, file)
// 	} else if homeId != "" {
// 		key = fmt.Sprintf("%s/home_%s_%d%s", config.Env, homeId, num, file)
// 	} else if name != "" {
// 		key = fmt.Sprintf("%s/%s_%d%s", config.Env, name, num, file)
// 	} else {
// 		key = fmt.Sprintf("%s/image/%s", config.Env, file)
// 	}

// 	ext := strings.ToLower(filepath.Ext(filename))

// 	if strings.HasPrefix(mediaType, "text/plain") {
// 		switch ext {
// 		case ".yaml", ".yml":
// 			mediaType = "application/x-yaml"
// 		case ".log":
// 			mediaType = "text/plain"
// 		}
// 	} else if mediaType == "application/x-ole-storage" || strings.HasPrefix(mediaType, "application/vnd") {
// 		switch ext {
// 		case ".doc", ".docx":
// 			mediaType = "application/doc"
// 		case ".xls", ".xlsx":
// 			mediaType = "application/xls"
// 		case ".ppt", ".pptx":
// 			mediaType = "application/ppt"
// 		}
// 	} else if mediaType == "application/x-rar-compressed" || mediaType == "application/vnd.rar" {
// 		mediaType = "application/x-rar-compressed"
// 	}

// 	_, err := f.r2Client.PutObject(ctx, &s3.PutObjectInput{
// 		Bucket:        aws.String(config.R2Bucket),
// 		Key:           aws.String(key),
// 		Body:          bytes.NewReader(mediaBytes),
// 		ContentType:   aws.String(mediaType),
// 		ContentLength: aws.Int64(int64(len(mediaBytes))),
// 	})

// 	if err != nil {
// 		return "", "", fmt.Errorf("failed to upload to R2: %w", err)
// 	}

// 	return filename, key, nil
// }

// func (f *File) RetrivedFile(ctx context.Context, key string) (string, error) {
// 	parts := strings.Split(key, "/")
// 	filename := parts[len(parts)-1]

// 	presignClient := s3.NewPresignClient(f.r2Client)

// 	presigned, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
// 		Bucket:                     aws.String(config.R2Bucket),
// 		Key:                        aws.String(key),
// 		ResponseContentDisposition: aws.String(fmt.Sprintf("attachment; filename=\"%s\"", filename)),
// 	}, s3.WithPresignExpires(24*time.Hour))
// 	if err != nil {
// 		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
// 	}

// 	return presigned.URL, nil
// }
